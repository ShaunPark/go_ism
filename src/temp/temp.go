package main

import (
	"encoding/json"
)

type Interface struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func main() {
	jsonStr := `{"id":"EAION", "name":"test interface"}`
	byt := []byte(jsonStr)

	println(string(byt))
	var inf Interface

	if err := json.Unmarshal(byt, &inf); err != nil {
		panic(err)
	}

	println(inf.Name)
	println(inf.Id)
}
