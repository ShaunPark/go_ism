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

	js, err = rule.GetField("FBFD000008")
	if err != nil {
		panic(err)
	}
	println("-------- Field ---------")
	println(js)

	js, err = rule.GetFieldGroup("FFGP000006")
	if err != nil {
		panic(err.Error())
	}
	println("-------- FieldGroup ---------")
	println(js)

	js, err = rule.GetService("FDBAS00002")
	if err != nil {
		panic(err.Error())
	}
	println("-------- Service ---------")
	println(js)

	js, err = rule.GetServiceMap("FMSGS00002")
	if err != nil {
		panic(err.Error())
	}
	println("-------- ServiceMap ---------")
	println(js)

	js, err = rule.GetServiceModel("SVCTMP0001")
	if err != nil {
		panic(err.Error())
	}
	println("-------- ServiceModel ---------")
	println(js)

	js, err = rule.GetServer("COM3")
	if err != nil {
		panic(err.Error())
	}
	println("-------- Server ---------")
	println(js)

	js, err = rule.GetSystem("FSYSEAI00001")
	if err != nil {
		panic(err.Error())
	}
	println("-------- System ---------")
	println(js)

}
