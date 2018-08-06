package transform

import "ism.com/common/rule/rmgr"

func Parse(msg []byte, dataId string) ([][][]string, error) {
	inDsrt := rmgr.GetDataStructure(dataId)
	var data [][][]byte
	var detail [][][][][]byte
	var rInfo [][]int

	if inDsrt != nil && inDsrt.Data != nil {

	}
	return nil, nil
}
