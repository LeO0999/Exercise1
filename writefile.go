package main

import (
	"fmt"
	"io/ioutil"
)

func WriteFile() {
	a := []byte(`[
		{
		  "name": "cofaxAdmin",
		  "class": "org.cofax.cds.AdminServlet"
		},
		{
		  "name": "fileServlet",
		  "class": "org.cofax.cds.FileServlet"
		}
	]
	`)
	err := ioutil.WriteFile("serve.json", a, 0777)

	if err != nil {
		fmt.Println(err)
	}
}
