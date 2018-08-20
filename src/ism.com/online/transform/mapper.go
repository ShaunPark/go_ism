package transform

import (
	"fmt"
	"strings"

	"ism.com/common/errorCode"
	"ism.com/common/ismerror"
	"ism.com/common/rule"
	"ism.com/common/rule/rmgr"
)

func mapping(input []ArrayInput, inDstr []rule.DataStructure, outDstr rule.DataStructure, sMap rule.ServiceMap, useTarget []bool) (ArrayInput, error) {

	var output ArrayInput
	data := outDstr.Data
	outData := make([][][]byte, len(data))
	outDetail := make([][][][][]byte, len(data))
	rCount := make([][]int, len(data))

	for i := 0; i < len(rCount); i++ {
		rCount[i] = make([]int, len(data[i].Detail))
	}

	for i := 0; i < len(rCount); i++ {
		for j := 0; j < len(rCount[i]); j++ {
			temp := data[i].Detail[j].RepeatCount
			if temp == 0 {
				rCount[i][j] = 1
			} else {
				rCount[i][j] = temp
			}
		}
	}

	for i := 0; i < len(data); i++ {
		if flg, err := rmgr.GetFieldGroup(data[i].MasterFieldGroupId); err != nil {
			return output, err
		} else {
			outData[i] = make([][]byte, len(flg.Fields))

			detail := data[i].Detail
			outDetail[i] = make([][][][]byte, len(detail))

			for j := 0; j < len(detail); j++ {
				if rCount[i][j] < 0 {
					rCount[i][j] = getRepeatCount(input, sMap.DetailMap[i][j])
				}
				outDetail[i][j] = make([][][]byte, rCount[i][j])

				for k := 0; k < rCount[i][j]; k++ {
					if detailFlg, err := rmgr.GetFieldGroup(detail[j].DetailFieldGroupId); err != nil {
						return output, err
					} else {
						outDetail[i][j][k] = make([][]byte, len(detailFlg.Fields))
					}
				}
			}
		}
	}

	for i := 0; i < len(outData); i++ {
		if flg, err := rmgr.GetFieldGroup(data[i].MasterFieldGroupId); err != nil {
			return output, err
		} else {
			fMap := flg.Fields
			if len(fMap) > 0 {
				if i >= len(sMap.MasterMap) {
					// DefaultLogger.logW("[warning] Master mapping rule. skip..");
					for x := 0; x < len(fMap); x++ {
						outData[i][x] = make([]byte, 0)
					}
				} else {
					if err := mapData(sMap.MasterMap[i], input, inDstr, outData[i], fMap, false, 0, useTarget); err != nil {
						return output, &ismerror.IsmError{-1, ""}
					}
				}
			}

			for j := 0; j < len(outDetail[i]); j++ {
				// 반복 횟수를 매핑에서 확정하도록 수정 2008-07-02
				if dFlg, err := rmgr.GetFieldGroup(outDstr.Data[i].Detail[j].DetailFieldGroupId); err == nil {
					fMap := dFlg.Fields
					if rCount[i][j] < 0 {
						rCount[i][j] = getRepeatCount(input, sMap.DetailMap[i][j])
					}
					for k := 0; k < rCount[i][j]; k++ {
						if len(fMap) > 0 {
							if err := mapData(sMap.DetailMap[i][j], input, inDstr, outDetail[i][j][k], fMap, false, k, useTarget); err != nil {
								return output, err
							}
						}
					}
				} else {
					return output, err
				}
			}
		}
	}
	output.RepeatInfo = rCount
	output.Data = outData
	output.Detail = outDetail

	return output, nil
}

func getRepeatCount(input []ArrayInput, dMap []rule.DataMap) int {
	rCount := 1
	for i := 0; i < len(dMap); i++ {
		srcs := dMap[i].Sources

		for j := 0; j < len(srcs); j++ {
			msgIndex := srcs[j].SourceMessageIndex
			if srcs[j].IsXml == "Y" {
				// return input[msgIndex].getXmlRepeatCount(srcs[j].getSourcePath());
			} else {
				dataIndex := srcs[j].SourceDataIndex
				detailIndex := srcs[j].SourceDetailIndex
				if dataIndex >= 0 && detailIndex >= 0 {
					inRptCnt := input[msgIndex].RepeatInfo
					return inRptCnt[dataIndex][detailIndex]
				}
			}
		}
	}
	return rCount
}

func mapData(maps []rule.DataMap, input []ArrayInput, inDstr []rule.DataStructure, buf [][]byte, fMap []rule.FieldMap, mapType bool, rCount int, useTarget []bool) error {
	fldCount := len(fMap)

	for i := 0; i < fldCount; i++ {
		cMap := maps[i]

		if mapType && cMap.Dataindex >= 0 || !mapType && cMap.Dataindex < 0 {
			continue
		}

		function := cMap.CustomFunction.String
		if function == "" || len(strings.TrimSpace(function)) <= 0 {
			buf[i] = getInput(cMap.Sources, input, rCount, useTarget)
		} else {
			// Object ret = convert(getInputs(cMap.getSource(), input, rCount, useTarget), function)
			// if (ret instanceof byte[]) {
			//     buf[i] = (byte[]) ret
			// } else {
			//     buf[i] = ((String) ret).getBytes()
			// }
		}
	}

	return nil
}

func getInput(sCol []rule.SourceColumn, input []ArrayInput, rCount int, useTarget []bool) []byte {
	var ret []byte
	src := make([][]byte, len(sCol))
	length := 0

	defaultValueIndex := -1

	for x := 0; x < len(sCol); x++ {
		if sCol[x].DefaultValue.Valid {
			defaultValueIndex = x
			break
		}
	}

	if defaultValueIndex != -1 {
		src = make([][]byte, 1)
		var err error
		if src[0], err = getSource(sCol[defaultValueIndex], input, rCount, useTarget); err != nil {

		}
		length += len(src[0])
	} else {
		for i := 0; i < len(sCol); i++ {
			var err error
			if src[i], err = getSource(sCol[i], input, rCount, useTarget); err != nil {

			}
			length += len(src[i])
		}
	}

	ret = make([]byte, length)
	cPos := 0
	for i := 0; i < len(src); i++ {
		src[i] = append(ret, src[i][cPos:len(src[i])]...)
		cPos += len(src[i])
	}

	return ret
}

func getSource(s rule.SourceColumn, input []ArrayInput, rCount int, useTarget []bool) ([]byte, error) {
	var ret []byte
	dataIndex := s.SourceDataIndex
	detailIndex := s.SourceDetailIndex
	columnIndex := s.SourceFieldIndex
	msgIndex := s.SourceMessageIndex

	// if (DefaultLogger.getLevel() >= Constants.LOG_TRIVIA) {
	//     DefaultLogger.logT("_D_ " + s.getFieldId() + ":" + dataIndex + ":" + detailIndex + ":" + columnIndex + ":" + msgIndex);
	// }

	if s.DefaultValue.Valid {
		ret = []byte(s.DefaultValue.String)
	} else {
		var inputBytes []byte

		if msgIndex == -1 || useTarget[msgIndex] {
			if msgIndex == -1 {
				msgIndex = len(input) - 1
			}

			if s.IsXml == "Y" {
				// ret = input[msgIndex].getXmlValue(s.getSourcePath(), rCount);
			} else {
				if dataIndex < 0 {
					if msgIndex >= len(input) || columnIndex >= len(input[msgIndex].Header) {
						return ret, &ismerror.IsmError{errorCode.TRNS_INVALID_MASTER_OFFSET_INDEX, fmt.Sprintf("failed to get mapping data[%d]:[%d]", msgIndex, columnIndex)}
					}
					inputBytes = input[msgIndex].Header[columnIndex]
				} else {
					if detailIndex < 0 {
						if msgIndex >= len(input) || dataIndex >= len(input[msgIndex].Data) || columnIndex >= len(input[msgIndex].Data[dataIndex]) {
							return ret, &ismerror.IsmError{errorCode.TRNS_INVALID_MASTER_OFFSET_INDEX, fmt.Sprintf("failed to get mapping data[%d]:[%d]:[%d]", msgIndex, dataIndex, columnIndex)}
						}
						inputBytes = input[msgIndex].Data[dataIndex][columnIndex]
					} else {
						// if (DefaultLogger.getLevel() >= Constants.LOG_TRIVIA) {
						//     DefaultLogger.logT("dataIndex = " + dataIndex + ", detailIndex = " + detailIndex + ", columnIndex = " + columnIndex);
						// }

						if msgIndex >= len(input) || dataIndex >= len(input[msgIndex].Detail) || detailIndex >= len(input[msgIndex].Detail[dataIndex]) ||
							rCount >= len(input[msgIndex].Detail[dataIndex][detailIndex]) || columnIndex >= len(input[msgIndex].Detail[dataIndex][detailIndex][rCount]) {
							return ret, &ismerror.IsmError{errorCode.TRNS_INVALID_DETAIL_OFFSET_INDEX, fmt.Sprintf("failed to get mapping data[%d]:[%d]:[%d]:[%d]:[%d]", msgIndex, dataIndex, detailIndex, rCount, columnIndex)}
						}

						inputBytes = input[msgIndex].Detail[dataIndex][detailIndex][rCount][columnIndex]
					}
				}
				/**
				 * @changed 2008-06-03
				 */
				if inputBytes != nil {
					ret = append([]byte{}, inputBytes[0:len(inputBytes)]...)
				}
			}
		}
	}

	return ret, nil
}
