package transform

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"strings"

	"ism.com/common/constants"
	"ism.com/common/errorCode"
	"ism.com/common/ismerror"
	"ism.com/common/rule"
	"ism.com/common/rule/rmgr"
)

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
	// var flds [][]byte

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
			mfgid := oDstrt.Data[i].MasterFieldGroupId

			if mfgid != "" {
				if mflg, err := rmgr.GetFieldGroup(mfgid); err != nil {
					return input, err
				} else {
					fMaps := mflg.Fields

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

					if flg, err := rmgr.GetFieldGroup(oDstrt.Data[i].Detail[j].DetailFieldGroupId); err != nil {
						//// TODO: error process
					} else {
						fMaps := flg.Fields
						rdelimiter := ""
						if oDstrt.Data[i].RecordDelimeter.Valid {
							rdelimiter = oDstrt.Data[i].RecordDelimeter.String
						}

						dfgid := oDstrt.Data[i].Detail[j].DetailFieldGroupId

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
		}
		input.DetailLength = detailLength
		input.Detail = detail
	}

	return input, nil
}

func getLength(input Input) int {
	length := 0

	// if input.Header != nil {
	// 	length += len(input.Header)
	// }
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
				var err error
				if tLen, err2 := findLengthInfo(lengths, dataIndex, detailIndex, i); err2 != nil {
					//// TODO: 에러처리
				} else {
					if fLen, err = getVariableLength(aInput, repeatCount, tLen, fMaps[i].FieldId, outDstrt); err != nil {
						//// TODO: 에러 처리
					}
				}
			}

			if fLen <= 0 {
				fLen = len(flds[i])
			}

			length += fLen
			temp := make([]byte, fLen)

			if len(flds[i]) >= fLen {
				if fld.FieldType == constants.TYPE_NUMBER && len(fld.FieldFormat) > 0 {
					var err error
					if temp, err = convertFormatNumber(flds[i], fld.FieldFormat, fLen, fld.Fillchar); err != nil {
						return []byte{}, &ismerror.IsmError{errorCode.TRNS_ARRAY_COPY_ERROR, "Number Format Error [" + fld.Name + "][" + fld.FieldFormat + "] input[" + string(flds[i]) + "]"}
					}
				} else {
					temp = append([]byte{}, flds[i][0:fLen]...)
				}
			} else if len(flds[i]) < fLen {
				if fld.FieldType == constants.TYPE_NUMBER && len(strings.TrimSpace(fld.FieldFormat)) > 0 {
					if temp, err = convertFormatNumber(flds[i], fld.FieldFormat, fLen, fld.Fillchar); err != nil {
						return []byte{}, &ismerror.IsmError{errorCode.TRNS_ARRAY_COPY_ERROR, "Number Format Error [" + fld.Name + "][" + fld.FieldFormat + "] input[" + string(flds[i]) + "]"}
					}
				} else {
					fChar := fld.Fillchar

					if fld.Aligntype == constants.ALIGN_LEFT {
						temp = append([]byte{}, flds[i][0:len(flds[i])]...)
						for j := len(flds[i]); j < len(temp); j++ {
							temp = append(temp, []byte(fChar)[:]...)
						}
					} else {
						j := 0
						temp = []byte{}

						for ; j < len(temp)-len(flds[i]); j++ {
							temp = append(temp, []byte(fChar)...)
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
		if fdLen >= 0 {
			row = append(row, flds[i][0:]...)
		}

		if fMaps[i].LengthFieldType.String != constants.LENGTH_NOT {
			for j := 0; j < len(lengthInfo); j++ {
				if dataIndex == -1 {
					if lengthInfo[j].dataIndex == dataIndex && lengthInfo[j].detailIndex == detailIndex && lengthInfo[j].fieldIndex == i {
						lengthInfo[j].offset = cPos
						lengthInfo[j].length = len(flds[i])
					}
				} else {
					if lengthInfo[j].dataIndex == dataIndex && lengthInfo[j].detailIndex == detailIndex && lengthInfo[j].repeatCount == repeatCount && lengthInfo[j].fieldIndex == i {
						lengthInfo[j].offset = cPos
						lengthInfo[j].length = len(flds[i])
					}
				}
			}
		}

		row = append(row, []byte(fdelimiter)...)
		cPos += len(flds[i])

		if i < (len(fMaps) - 1) {
			cPos += len([]byte(fdelimiter))
		}

	}
	// bPos := cPos
	row = append(row, []byte(rdelimeter)...)
	cPos += rLen

	// aPos := cPos
	return row, nil
}

func getVariableLength(input ArrayInput, repeatCount int, lengthInfo rule.LengthFieldInfo, fId string, outDstrt rule.DataStructure) (int, error) {
	length := 0
	var data [][]byte
	dataIndex := lengthInfo.LengthDataIndex
	detailIndex := lengthInfo.LengthDetailIndex
	columnIndex := lengthInfo.LengthFieldIndex

	if dataIndex < 0 {
		// data = input.getHeader()
	} else {
		if detailIndex < 0 {
			if input.Data == nil {
				return -1, &ismerror.IsmError{errorCode.TRNS_MAPPED_DATA_NULL, ""}
				// 가변길이 필드 길이 계산 [%S]- 매핑된 결과에 있는 데이터가 Null 입니다.
				// int eCode = ErrorCode.TRNS_MAPPED_DATA_NULL; // TRNS 54
				// String[] params = ExUtil.params(fId, dataIndex);
				// throw new TransformerException(eCode, ErrorMessage.getMessage(eCode, params), params);
			} else if len(input.Data) < dataIndex {
				return -1, &ismerror.IsmError{errorCode.TRNS_MAPPED_DATA_LENGTH_ERROR, ""}
				// 가변길이 필드 길이 계산 [%S]- 매핑된 결과에 있는 데이터의 갯수[%S]가 가변길이 필드에 정의된 데이터 인덱스[%S]보다 작거나 같습니다.
				// int eCode = ErrorCode.TRNS_MAPPED_DATA_LENGTH_ERROR;  // TRNS 55
				// String[] params = ExUtil.params(fId, input.getData().length, dataIndex);
				// throw new TransformerException(eCode, ErrorMessage.getMessage(eCode, params), params);
			}
			data = input.Data[dataIndex]
		} else {
			if input.Detail == nil {
				return -1, &ismerror.IsmError{errorCode.TRNS_MAPPED_DETAIL_NULL, ""}
				// // 가변길이 필드 길이 계산 [%S]- 매핑된 결과에 있는 디테일이 Null 입니다.
				// int eCode = ErrorCode.TRNS_MAPPED_DETAIL_NULL; // TRNS 56
				// String[] params = ExUtil.params(fId);
				// throw new TransformerException(eCode, ErrorMessage.getMessage(eCode, params), params);
			} else if len(input.Detail) < dataIndex {
				return -1, &ismerror.IsmError{errorCode.TRNS_MAPPED_DATA_LENGTH_ERROR, ""}
				// // 가변길이 필드 길이 계산 [%S]- 매핑된 결과에 있는 데이터의 갯수[%S]가 가변길이 필드에 정의된 데이터 인덱스[%S]보다 작거나 같습니다.
				// int eCode = ErrorCode.TRNS_MAPPED_DATA_LENGTH_ERROR; // TRNS 55
				// String[] params = ExUtil.params(fId, input.getDetail().length, dataIndex);
				// throw new TransformerException(eCode, ErrorMessage.getMessage(eCode, params), params);
			} else if len(input.Detail[dataIndex]) < detailIndex {
				return -1, &ismerror.IsmError{errorCode.TRNS_MAPPED_DETAIL_LENGTH_ERROR, ""}
				// // 가변길이 필드 길이 계산 [%S]- 매핑된 결과에 있는 디테일의 갯수[%S]가 가변길이 필드에 정의된 디테일 인덱스[%S]보다 작거나 같습니다.
				// int eCode = ErrorCode.TRNS_MAPPED_DETAIL_LENGTH_ERROR; // TRNS 57
				// String[] params = ExUtil.params(fId, input.getDetail()[dataIndex].length, detailIndex);
				// throw new TransformerException(eCode, ErrorMessage.getMessage(eCode, params), params);
			}
			data = input.Detail[dataIndex][detailIndex][repeatCount]
		}
	}

	if len(data) < columnIndex {
		return -1, &ismerror.IsmError{errorCode.TRNS_MAPPED_FIELD_LENGTH_ERROR, ""}
		// // 가변길이 필드 길이 계산 [%S]- 매핑된 결과에 있는 필드의 갯수[%S]가 가변길이 필드에 정의된 필드 인덱스[%S]보다 작거나 같습니다.
		// int eCode = ErrorCode.TRNS_MAPPED_FIELD_LENGTH_ERROR; // TRNS 58
		// String[] params = ExUtil.params(fId, data.length, columnIndex);
		// throw new TransformerException(eCode, ErrorMessage.getMessage(eCode, params), params);
	} else if data[columnIndex] == nil {
		return -1, &ismerror.IsmError{errorCode.TRNS_VARIABLE_LENGTH_FIELD_VALUE_NULL, ""}
		//
		// // 가변길이 필드 길이 계산 [%S]- 매핑된 결과에 있는 필드의 값이 Null 입니다. DataIndex[%S], DetailIndex[%S], FieldIndex[%S]
		// String[] params = ExUtil.params(fId, dataIndex, detailIndex, columnIndex);
		// throw new TransformerException(ErrorCode.TRNS_VARIABLE_LENGTH_FIELD_VALUE_NULL, ErrorMessage.getMessage(
		//         ErrorCode.TRNS_VARIABLE_LENGTH_FIELD_VALUE_NULL, params), params);
	}

	var flg rule.FieldGroup
	var err error
	if detailIndex < 0 {
		if flg, err = rmgr.GetFieldGroup(outDstrt.Data[dataIndex].MasterFieldGroupId); err != nil {
			return -1, err
		}
	} else {
		if flg, err = rmgr.GetFieldGroup(outDstrt.Data[dataIndex].Detail[detailIndex].DetailFieldGroupId); err != nil {
			return -1, err
		}
	}

	if fld, err := rmgr.GetField(flg.Fields[columnIndex].FieldId); err != nil {
		// TODO:
		return -1, err
	} else {
		if fld.FieldType == constants.TYPE_BINARY {
			buf := bytes.NewReader(data[columnIndex])
			err := binary.Read(buf, binary.LittleEndian, &length)
			if err != nil {
				return -1, &ismerror.IsmError{errorCode.TRNS_LENGTH_PARSING_ERROR, ""}
			}

			var rErr error
			switch fld.FieldLength {
			case 1:
				var temp int8
				rErr = binary.Read(buf, binary.LittleEndian, &temp)
				length = int(temp)
			case 2:
				var temp int16
				rErr = binary.Read(buf, binary.LittleEndian, &temp)
				length = int(temp)
			case 4:
				var temp int32
				rErr = binary.Read(buf, binary.LittleEndian, &temp)
				length = int(temp)
			case 8:
				var temp int64
				rErr = binary.Read(buf, binary.LittleEndian, &temp)
				length = int(temp)
			default:
				rErr = &ismerror.IsmError{errorCode.TRNS_LENGTH_PARSING_ERROR, "Unsupported field length. 2 and 4 are available."}
			}
			if rErr != nil {
				return -1, rErr
			}
		} else if fld.FieldType == constants.TYPE_DATE {
			return -1, &ismerror.IsmError{errorCode.TRNS_LENGTH_PARSING_ERROR, "Date format can not used for length field."}
			// int eCode = ErrorCode.TRNS_LENGTH_PARSING_ERROR;
			// TransformerException te = new TransformerException(eCode, "Date format can not used for length field :" + ErrorMessage.getMessage(eCode));
			// throw te;
		} else {
			if length, err = strconv.Atoi(strings.TrimSpace(string(data[columnIndex]))); err != nil {
				return -1, &ismerror.IsmError{errorCode.TRNS_LENGTH_PARSING_ERROR, ""}
			}
		}

		length -= lengthInfo.DiffValue

		if length <= 0 {
			return -1, &ismerror.IsmError{errorCode.TRNS_LENGTH_PARSING_ERROR, ""}
		}
		return length, nil
	}
}

func convertFormatNumber(input []byte, format string, length int, fillChar string) ([]byte, error) {
	// if(DefaultLogger.getLevel() >= Constants.LOG_VERBOSE) {
	//     DefaultLogger.logV("convert format numner [{}][{}][{}]", new String(input), format, length );
	// }
	var rtnValue []byte
	// String roundingMode = Common.getProperty("transformer.rounding.mode");
	// rtnValue = new byte[length];
	// String inValue = new String(input).trim();
	// if(inValue.isEmpty()) {
	//     // throw new TransformerException(ErrorCode.TRNS_NO_DATA, "input data is empty.");
	//   // 2013-12-17 송신 데이터가 없을경우 Number 형 변환 시 0 으로 기본값 설정.
	//   for(int i=0; i < rtnValue.length; i++) {
	//     rtnValue[i] = fillChar;
	//   }
	//   return rtnValue;
	// }
	// BigDecimal value = new BigDecimal(inValue);
	// byte[] trandata = null;
	// if(!format.startsWith("%")) {
	//     DecimalFormat formatter = new DecimalFormat(format);
	//     if(roundingMode == null || roundingMode.length() == 0 ) {
	//         formatter.setRoundingMode(RoundingMode.DOWN);
	//     } else if(roundingMode.equalsIgnoreCase("HALF_EVEN")) {
	//         formatter.setRoundingMode(RoundingMode.HALF_EVEN);
	//     } else if(roundingMode.equalsIgnoreCase("HALF_DOWN")) {
	//         formatter.setRoundingMode(RoundingMode.HALF_DOWN);
	//     } else if(roundingMode.equalsIgnoreCase("HALF_UP")) {
	//         formatter.setRoundingMode(RoundingMode.HALF_UP);
	//     } else if(roundingMode.equalsIgnoreCase("UP")) {
	//         formatter.setRoundingMode(RoundingMode.UP);
	//     } else if(roundingMode.equalsIgnoreCase("DOWN")) {
	//         formatter.setRoundingMode(RoundingMode.DOWN);
	//     } else if(roundingMode.equalsIgnoreCase("CEILING")) {
	//         formatter.setRoundingMode(RoundingMode.CEILING);
	//     } else if(roundingMode.equalsIgnoreCase("FLOOR")) {
	//         formatter.setRoundingMode(RoundingMode.FLOOR);
	//     }
	//     trandata = formatter.format(value).toString().getBytes();
	// } else
	// if(format.endsWith("f")) { // decimal integer
	//     int index = format.lastIndexOf(".");
	//     if(index > 0) { // % 로 시작하여야 하니 0 이 아닌 1부터
	//         int scale = Integer.parseInt(format.substring(index + 1, format.length() - 1));
	//         if(roundingMode == null || roundingMode.length() == 0 ) {
	//             value = value.setScale(scale, RoundingMode.DOWN);
	//         } else if(roundingMode.equalsIgnoreCase("HALF_EVEN")) {
	//             value = value.setScale(scale, RoundingMode.HALF_EVEN);
	//         } else if(roundingMode.equalsIgnoreCase("HALF_DOWN")) {
	//             value = value.setScale(scale, RoundingMode.HALF_DOWN);
	//         } else if(roundingMode.equalsIgnoreCase("HALF_UP")) {
	//             value = value.setScale(scale, RoundingMode.HALF_UP);
	//         } else if(roundingMode.equalsIgnoreCase("UP")) {
	//             value = value.setScale(scale, RoundingMode.UP);
	//         } else if(roundingMode.equalsIgnoreCase("DOWN")) {
	//             value = value.setScale(scale, RoundingMode.DOWN);
	//         } else if(roundingMode.equalsIgnoreCase("CEILING")) {
	//             value = value.setScale(scale, RoundingMode.CEILING);
	//         } else if(roundingMode.equalsIgnoreCase("FLOOR")) {
	//             value = value.setScale(scale, RoundingMode.FLOOR);
	//         }
	//     }
	//     trandata = new Formatter().format(format, value).toString().trim().getBytes();
	// } else {
	//     trandata = new Formatter().format(format, value.longValue()).toString().trim().getBytes();
	// }
	// if(length > trandata.length) {
	//     int i = 0;
	//     for(; i<(length - trandata.length); i++) {
	//         rtnValue[i] = fillChar;
	//     }
	//     System.arraycopy(trandata, 0, rtnValue, i, trandata.length);
	// } else {
	//     if(trandata[0] == '-') {
	//         rtnValue[0] = '-';
	//         System.arraycopy(trandata, trandata.length - length + 1, rtnValue, 1, length - 1);
	//     } else {
	//         System.arraycopy(trandata, trandata.length - length, rtnValue, 0, length);
	//     }
	// }
	// if(DefaultLogger.getLevel() >= Constants.LOG_VERBOSE) {
	//     DefaultLogger.logV("convert Format Number [{}] in[{}] out[{}]", format, new String(input), new String(rtnValue) );
	// }
	return rtnValue, nil
}
