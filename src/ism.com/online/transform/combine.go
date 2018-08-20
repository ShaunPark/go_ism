package transform

import (
	"strings"

	"ism.com/common/constants"
	"ism.com/common/errorCode"
	"ism.com/common/ismerror"
	"ism.com/common/rule"
	"ism.com/common/rule/rmgr"
)

type Length struct {
	dataIndex   int    //dataIndex // -1
	detailIndex int    //= -1;
	fieldIndex  int    //= -1;
	repeatCount int    //= -1;
	diffValue   int    //= 0;
	fieldId     string //;
	offset      int    // = -1;
	length      int    //= -1;
	lenType     int    //= Constants.TOTAL_LENGTH;
}

type Output struct {
	byteMessage  []byte
	input        Input
	dataOffset   []int
	detailOffset [][][]int
}

func combine(aInput ArrayInput, oDstrt rule.DataStructure, length []Length) (Output, error) {
	// if ( oDstrt.getMessageType() == Constants.MESSAGE_ISO) {
	//     return combineISO(aInput, oDstrt);
	//
	// }

	if input, err := makeArrayInputToInput(aInput, oDstrt, length); err != nil {
		return Output{}, err
	} else {

		msg := make([]byte, getLength(input))
		cPos := 0
		var dataOffset []int
		var detailOffset [][][]int

		if input.Data != nil {
			dataOffset = make([]int, len(input.Data))
			detailOffset = make([][][]int, len(input.Data))

			for i := 0; i < len(input.Data); i++ {
				if input.Data[i] != nil {
					msg = append(msg, input.Data[i]...)
					dataOffset[i] = cPos
					cPos += len(input.Data[i])
				}
				if input.Detail[i] != nil {
					detailOffset[i] = make([][]int, len(input.Detail[i]))
					for j := 0; j < len(input.Detail[i]); j++ {
						if input.Detail[i][j] != nil {
							detailOffset[i][j] = make([]int, len(input.Detail[i][j]))
							for k := 0; k < len(input.Detail[i][j]); k++ {
								msg = append(msg, input.Detail[i][j][k]...)
								//System.arraycopy(input.Detail[i][j][k], 0, msg, cPos, input.getDetail()[i][j][k].length)
								detailOffset[i][j][k] = cPos
								cPos += len(input.Detail[i][j][k])
							}
						}
					}
				}
			}
		}

		return Output{msg, input, dataOffset, detailOffset}, nil
	}
}

func makeArrayInputToInput(aInput ArrayInput, oDstrt rule.DataStructure, length []Length) (Input, error) {
	var input Input
	repeatCount := aInput.RepeatInfo
	var flds [][]byte

	if aInput.Data != nil {
		data := make([][]byte, len(aInput.Data))
		dataLength := make([]int, len(aInput.Data))
		for i := 0; i < len(aInput.Data); i++ {
			flds := aInput.Data[i]
			rdelimiter := ""
			if oDstrt.Data[i].RecordDelimeter.Valid {
				rdelimiter = oDstrt.Data[i].RecordDelimeter.String
			}
			// 2008-07-02 마스터가 없을 수 있으므로 수정
			if oDstrt.Data[i].Master != nil {
				fMaps := oDstrt.Data[i].Master.Fields
				mfgid := oDstrt.Data[i].MasterFieldGroupId

				if mfg, err := rmgr.GetFieldGroup(mfgid); err != nil {
					return input, &ismerror.IsmError{-1, ""}
				} else {
					fdelimiter := ""
					if mfg.FieldDelimeter.Valid {
						fdelimiter = mfg.FieldDelimeter.String
					}
					var err error
					if data[i], err = makeRow(aInput, flds, fMaps, length, i, -1, -1, oDstrt.Lengths, oDstrt, rdelimiter, fdelimiter); err != nil {

					} else {
						dataLength[i] = len(data[i])
					}
				}
			} else {
				dataLength[i] = 0
			}
		}
		input.DataLength = dataLength
		input.Data = data
	}

	if aInput.Detail != nil {
		detail := make([][][][]byte, len(aInput.Data))
		detailLength := make([][][]int, len(aInput.Data))

		for i := 0; i < len(aInput.Data); i++ {
			detailLength[i] = make([][]int, len(aInput.Detail[i]))
			detail[i] = make([][][]byte, len(aInput.Detail[i]))

			for j := 0; j < len(aInput.Detail[i]); j++ {
				detail[i][j] = make([][]byte, repeatCount[i][j])
				detailLength[i][j] = make([]int, repeatCount[i][j])
				for k := 0; k < repeatCount[i][j]; k++ {
					flds := aInput.Detail[i][j][k]
					fMaps := oDstrt.Data[i].Detail[j].Detail.Fields
					rdelimiter := ""
					if oDstrt.Data[i].RecordDelimeter.Valid {
						rdelimiter = oDstrt.Data[i].RecordDelimeter.String
					}

					dfgid := oDstrt.Data[i].Detail[j].DetailId

					if dfg, err := rmgr.GetFieldGroup(dfgid); err != nil {
						return input, err
					} else {
						fdelimiter := ""
						if dfg.FieldDelimeter.Valid {
							fdelimiter = dfg.FieldDelimeter.String
						}
						var err2 error
						if detail[i][j][k], err2 = makeRow(aInput, flds, fMaps, length, i, j, k, oDstrt.Lengths, oDstrt, rdelimiter, fdelimiter); err2 != nil {
							return input, err2
						} else {
							detailLength[i][j][k] = len(detail[i][j][k])
						}
					}
				}
			}
		}
		input.DetailLength = detailLength
		input.Detail = detail
	}

	return input, nil
}

func getLength(input Input) int {
	length := 0

	if input.Header != nil {
		length += len(input.Header)
	}
	for i := 0; input.Data != nil && i < len(input.Data); i++ {
		if input.Data[i] != nil {
			length += len(input.Data[i])
		}
	}
	for i := 0; input.Detail != nil && i < len(input.Detail); i++ {
		for j := 0; input.Detail[i] != nil && j < len(input.Detail[i]); j++ {
			for k := 0; input.Detail[i][j] != nil && k < len(input.Detail[i][j]); k++ {
				if input.Detail[i][j][k] != nil {
					length += len(input.Detail[i][j][k])
				}
			}
		}
	}

	return length
}

func makeRow(aInput ArrayInput, flds [][]byte, fMaps []rule.FieldMap, lengthInfo []Length, dataIndex int, detailIndex int, repeatCount int, lengths []rule.LengthFieldInfo, outDstrt rule.DataStructure, rdelimeter string, fdelimiter string) ([]byte, error) {
	length := 0

	rLen := len([]byte(rdelimeter))
	fdLen := len([]byte(fdelimiter))

	for i := 0; i < len(fMaps); i++ {

		if fld, err := rmgr.GetField(fMaps[i].FieldId); err != nil {

		} else {
			fLen := fld.FieldLength
			if fLen == 0 {
				fLen = getVariableLength(aInput, repeatCount, findLengthInfo(lengths, dataIndex, detailIndex, i), fMaps[i].FieldId, outDstrt)
			}

			if fLen <= 0 {
				fLen = len(flds[i])
			}

			length += fLen
			temp := make([]byte, fLen)

			if len(flds[i]) >= fLen {
				if fld.FieldType == constants.TYPE_NUMBER && len(fld.FieldFormat) > 0 {
					var err error
					if temp, err = convertFormatNumber(flds[i], fld.FieldFormat, fLen, fld.FillChar); err != nil {
						return []byte{}, &ismerror.IsmError{errorCode.TRNS_ARRAY_COPY_ERROR, "Number Format Error [" + fld.Name + "][" + fld.FieldFormat + "] input[" + string(flds[i]) + "]"}
					}
				} else {
					temp = append([]byte{}, flds[i][0:fLen]...)
				}
			} else if len(flds[i]) < fLen {
				if fld.FieldType == constants.TYPE_NUMBER && fld.FieldFormat != nil && len(strings.TrimSpace(fld.FieldFormat())) > 0 {
					if temp, err = convertFormatNumber(flds[i], fld.FieldFormat, fLen, fld.FillChar); err != nil {
						return []byte{}, &ismerror.IsmError{errorCode.TRNS_ARRAY_COPY_ERROR, "Number Format Error [" + fld.Name + "][" + fld.FieldFormat + "] input[" + string(flds[i]) + "]"}
					}
				} else {
					fChar := fld.FillChar

					if fld.AlignType == Constants.ALIGN_LEFT {
						temp = append([]byte{}, flds[i][0:len(flds[i])]...)
						for j := flds[i].length; j < temp.length; j++ {
							temp = append(temp, []byte(fChar))
						}
					} else {
						j := 0
						temp = []byte{}

						for ; j < len(temp)-len(flds[i]); j++ {
							temp = append(temp, []byte(fChar))
						}
						temp = append(temp, flds[i][0:len(flds[i])]...)
					}
				}
			}
			flds[i] = temp
		}
	}

	row := make([]byte, length+fdLen*(len(fMaps)-1)+rLen)

	cPos := 0
	for i := 0; i < len(fMaps); i++ {
		if fdLen > 0 {
			System.arraycopy(flds[i], 0, row, cPos, flds[i].length)
		} else {
			System.arraycopy(flds[i], 0, row, cPos, flds[i].length)
		}
		if fMaps[i].getLengthFieldType() != Constants.LENGTH_NOT {
			for j := 0; j < len(lengthInfo); j++ {
				if dataIndex == -1 {
					if lengthInfo[j].getDataIndex() == dataIndex && lengthInfo[j].getDetailIndex() == detailIndex && lengthInfo[j].getFieldIndex() == i {
						lengthInfo[j].setOffset(cPos)
						lengthInfo[j].setLength(flds[i].length)
					}
				} else {
					if lengthInfo[j].getDataIndex() == dataIndex && lengthInfo[j].getDetailIndex() == detailIndex && lengthInfo[j].getRepeatCount() == repeatCount && lengthInfo[j].getFieldIndex() == i {
						lengthInfo[j].setOffset(cPos)
						lengthInfo[j].setLength(flds[i].length)
					}
				}
			}
		}
		System.arraycopy(fdelimiter.getBytes(), 0, row, cPos, fdelimiter.getBytes().length)
		cPos += flds[i].length
		if i < (fMaps.length - 1) {
			cPos += fdelimiter.getBytes().length
		}

	}
	bPos := cPos
	System.arraycopy(rdelimeter.getBytes(), 0, row, cPos, rLen)

	cPos += rLen

	aPos := cPos

	DefaultLogger.logV(" bpos = " + bPos + " apos = " + aPos)
	return row
}
