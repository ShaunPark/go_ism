package main

import (
	"fmt"

	"ism.com/online/transform"
)

type Interface struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func main() {
	// js, err := rule.GetInterface("CBSBN0001002")
	// if err != nil {
	// 	panic(err)
	// }
	// println("--------- Interface --------")
	// println(js)
	//
	// js, err = rule.GetApplication("A018")
	// if err != nil {
	// 	panic(err)
	// }
	// println("--------- Application --------")
	// println(js)
	//
	// js, err = rule.GetDataStructure("FDST000008")
	// if err != nil {
	// 	panic(err)
	// }
	// println("-------- DataStructure ---------")
	// println(js)
	//
	// js, err = rule.GetField("FBFD000008")
	// if err != nil {
	// 	panic(err)
	// }
	// println("-------- Field ---------")
	// println(js)
	//
	// js, err = rule.GetFieldGroup("FFGP000006")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// println("-------- FieldGroup ---------")
	// println(js)
	//
	// js, err = rule.GetService("FDBAS00002")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// println("-------- Service ---------")
	// println(js)
	//
	// js, err = rule.GetServiceMap("SMAP000014")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// println("-------- ServiceMap ---------")
	// println(js)
	//
	// js, err = rule.GetServiceModel("SVCTMP0001")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// println("-------- ServiceModel ---------")
	// println(js)
	//
	// js, err = rule.GetServer("COM3")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// println("-------- Server ---------")
	// println(js)
	//
	// js, err = rule.GetSystem("FSYSEAI00001")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// println("-------- System ---------")
	// println(js)

	if output, err := transform.Parse([]byte("12345678901234567890asdfgtrewqasdfgtrewq000000021234567890123456789000000020asdfgtrewqasdfgtrewq1234567890123456789000000010asdfgtrewq1234567890123456789000000015asdfgtrewqqwert"), "FDST000665"); err != nil {
		panic(err)
	} else {
		for _, d := range output.Data {
			for _, dd := range d {
				fmt.Println(string(dd))
			}
		}
		fmt.Println(output.Detail)

		for _, d := range output.Detail {
			for _, dd := range d {
				for _, ddd := range dd {
					for _, dddd := range ddd {
						fmt.Println(string(dddd))
					}
				}
			}
		}
	}
	// ByteHeader []byte
	// ByteData   []byte
	// Header     [][]byte
	// Data       [][][]byte
	// Detail     [][][][][]byte
	// RepeatInfo [][]int
}
