package main

import (
	"ism.com/common/rule"
)

type Interface struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func main() {
	js, err := rule.GetInterface("CUSBN0003022")
	if err != nil {
		panic(err)
	}
	println(js)
}
