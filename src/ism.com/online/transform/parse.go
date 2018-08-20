package transform

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"ism.com/common/errorCode"
	"ism.com/common/ismerror"
	"ism.com/common/rule"
	"ism.com/common/rule/rmgr"
)

func Parse(msg []byte, dataId string) (ArrayInput, error) {
	inDsrt := rmgr.GetDataStructure(dataId)
	var input ArrayInput

	cPos := 0
	dataInf := inDsrt.Data
	data := make([][][]byte, len(dataInf))
	detail := make([][][][][]byte, len(dataInf))
	rInfo := make([][]int, len(dataInf))
	input.Data = data
	input.Detail = detail
	input.RepeatInfo = rInfo

	for j := 0; j < len(dataInf); j++ {
		if flgGrp, err := rmgr.GetFieldGroup(dataInf[j].MasterFieldGroupId); err != nil {
			return input, err
		} else {
			var recordDel []byte
			rd := dataInf[j].RecordDelimeter.String
			if (rd != "" && rd != "null") && (len(rd) > 0 || rd == "\n") {
				recordDel = []byte(dataInf[j].RecordDelimeter.String)
			}

			data[j] = make([][]byte, len(flgGrp.Fields))
			var err error
			if cPos, err = parseRow(cPos, msg, input, flgGrp, data[j], j, -1, -1, recordDel, (dataInf[j].Detail == nil || len(dataInf[j].Detail) == 0), inDsrt.Lengths, inDsrt); err != nil {
				return input, err
			}

			detailInf := dataInf[j].Detail
			detail[j] = make([][][][]byte, len(detailInf))
			rInfo[j] = make([]int, len(detailInf))

			// Detail parsing
			for k := 0; k < len(detailInf); k++ {
				if flgDetail, err := rmgr.GetFieldGroup(detailInf[k].DetailFieldGroupId); err != nil {

				} else {
					// detail repeat 처리
					rCount := detailInf[k].RepeatCount

					// if (DefaultLogger.getLevel() >= Constants.LOG_VERBOSE) {
					// 	if (fldGroupDetail != null) {
					// 				DefaultLogger.logV("fldGroupDetail : {}", fldGroupDetail.getId());
					// 		} else {
					// 				DefaultLogger.logV("fldGroupDetail is null.");
					// 		}
					// }
					if rCount <= 0 {
						rDIdx := detailInf[k].RepeatCountDataIndex
						rFIdx := detailInf[k].RepeatCountFieldIndex
						// if (DefaultLogger.getLevel() >= Constants.LOG_VERBOSE) {
						// 		DefaultLogger.logV("repeat dataIndex : {}  fieldIndex : {}", repeatDataIndex, repeatFieldIndex);
						// }
						if rDIdx < 0 {
							rCount = -1
						} else {
							// if (DefaultLogger.getLevel() >= Constants.LOG_VERBOSE) {
							// 		DefaultLogger.logV("input.getData().length ? {}", input.getData().length);
							// }
							if input.Data == nil || len(input.Data) < rDIdx {
								return input, &ismerror.IsmError{errorCode.TRNS_INVALID_REPEAT_COUNT_DATA_INDEX, ""}
							}
							if len(input.Data[rDIdx]) < rFIdx {
								return input, &ismerror.IsmError{errorCode.TRNS_INVALID_REPEAT_COUNT_FIELD_INDEX, ""}
							}
							var err error
							if rCount, err = strconv.Atoi(string(input.Data[rDIdx][rFIdx])); err != nil {
								return input, &ismerror.IsmError{errorCode.TRNS_INVALID_REPEAT_COUNT_FORMAT, ""}
							}
						}
					}

					/**
					 * @changed ???
					 *          from '<= 0' to '< 0'
					 *          for online db pattern.
					 *          it's possible result data of retrieval is null
					 *          and repeat count is set to 0
					 */
					// if rCount < 0 {
					// 		if recordDel != nil && len(recordDel) > 0 {
					// 				ByteArray ba = new ByteArray(in, cPos);
					// 				repeatCount = ba.tokenize(recordDel);
					// 		} else {
					// 			return -1, &ismerror.IsmError{ErrorCode.TRNS_INVALID_REPEAT_COUNT_VALUE,""}
					// 		}
					// }

					detail[j][k] = make([][][]byte, rCount)
					rInfo[j][k] = rCount

					for rc := 0; rc < rCount; rc++ {
						if flgDetail.Fields != nil {
							detail[j][k][rc] = make([][]byte, len(flgDetail.Fields))
							if cPos, err = parseRow(cPos, msg, input, flgDetail, detail[j][k][rc], j, k, rc, recordDel, (len(detailInf) == k+1) && (rc+1 == rCount), inDsrt.Lengths, inDsrt); err != nil {
								return input, err
							}
						}
					}
				}
			}
		}
	}
	return input, nil
}

func parseRow(cPos int, in []byte, input ArrayInput, flg rule.FieldGroup, tempArray [][]byte, dIdx int, dtIdx int, rIdx int, rd []byte, isLast bool, lengths []rule.LengthFieldInfo, inDstrt rule.DataStructure) (int, error) {

	fldDelimeter := strings.TrimSpace(flg.FieldDelimeter.String)

	if len(fldDelimeter) > 0 || flg.FieldDelimeter.String == "\n" {
		//return parseRowDel(cPos, in, input, fldGroup, tempArray, repeatIndex, recordDelimeter, lastDetail);
	}
	fmap := flg.Fields
	fldCnt := len(fmap)

	for i := 0; i < fldCnt; i++ {
		var fld rule.Field
		var err error
		if fld, err = rmgr.GetField(fmap[i].FieldId); err != nil {
			panic(err)
		}

		if fld.FieldLength == 0 {
			if lengthInfo, err := findLengthInfo(lengths, dIdx, dtIdx, i); err != nil {
				panic(err)
			} else {
				if vLen, err := getVariableLength(input, rIdx, lengthInfo, fmap[i].FieldId, inDstrt); err != nil {
					return -1, err
				} else {
					tempArray[i] = make([]byte, vLen)
				}

			}
		} else {
			tempArray[i] = make([]byte, fld.FieldLength)
		}

		// try {
		copyLength := len(tempArray[i])

		if (len(in) - cPos) < len(tempArray[i]) {
			copyLength = len(in) - cPos

		}

		if cPos > len(in) || cPos+copyLength > len(in) {
			return -1, &ismerror.IsmError{errorCode.TRNS_ARRAY_COPY_ERROR, ""}
		}
		tempArray[i] = append([]byte{}, in[cPos:(cPos+copyLength)]...)

		cPos += len(tempArray[i])
	}

	return cPos, nil
}

func findLengthInfo(lengths []rule.LengthFieldInfo, dIdx int, dtIdx int, fIdx int) (rule.LengthFieldInfo, error) {
	for i := 0; i < len(lengths); i++ {
		if lengths[i].DataIndex == dIdx && lengths[i].DetailIndex == dtIdx && lengths[i].FieldIndex == fIdx {
			return lengths[i], nil
		}
	}
	var ret rule.LengthFieldInfo
	return ret, errors.New(fmt.Sprint("Length field for (", dIdx, ",", dtIdx, ",", fIdx, ") is not defined"))
}
