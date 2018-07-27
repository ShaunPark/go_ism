package myhttp

import (
	"time"
)

type MyhttpClient struct {
}

var idx int

func (client *MyhttpClient) Call() string {
	idx += 1
	println("called ", idx)
	localIdx := idx
	if idx%2 == 0 {
		time.Sleep(1 * time.Second)
	} else {
		time.Sleep(2 * time.Second)
	}
	println("end of ", localIdx)
	return "call success!!!"
}
