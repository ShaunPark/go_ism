package main

import (
	"ism.com/common/rule"
)

type Interface struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func main() {
	js, err := rule.GetInterface("CBSBN0001002")
	if err != nil {
		panic(err)
	}
	println("--------- Interface --------")
	println(js)

	js, err = rule.GetApplication("A018")
	if err != nil {
		panic(err)
	}
	println("--------- Application --------")
	println(js)

	js, err = rule.GetDataStructure("FDST000008")
	if err != nil {
		panic(err)
	}
	println("-------- DataStructure ---------")
	println(js)
}
