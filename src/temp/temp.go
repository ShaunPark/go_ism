package main

import (
	"ism.com/common/rule"
)

type Interface struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func main() {
	js, err := rule.GetField("FBFD000003")
	if err != nil {
		panic(err)
	}
	println(js)
}
