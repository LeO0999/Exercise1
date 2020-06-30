package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Readfile()  {
	file, err := ioutil.ReadFile("serve.json")
	if err != nil {
		fmt.Println(err)
	}

	data := make([]Serve, 2)

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data); i++ {
		fmt.Println("Name: ", data[i].Name)
		fmt.Println("Class: ", data[i].Class)
	}
}