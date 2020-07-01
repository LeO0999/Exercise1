package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Serve struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

func main() {

	//Bai 1: Ghi file
	//WriteFile()
	// Bai 2: đọc file
	file, err := ioutil.ReadFile("serve.json")
	if err != nil {
		fmt.Println(err)
	}

	data := make([]Serve, 2)

	_ = json.Unmarshal(file, &data)

	for _, i := range data {
		log.Println("Name: ", i.Name)
		log.Println("Class: ", i.Class)
	}

	// Bai 3: lọc kết quả với class có chứa từ `admin` in ra đối tượng phù hợp
	xet := 0
	for _, i := range data {

		if strings.Contains(i.Class, "admin") {
			fmt.Println("Doi tuong co chu 'admin'la ")
			fmt.Println(i)
			xet += 1
		}
	}
	if xet <= 0 {
		fmt.Println("Ko co doi tuong nao co class chua 'admin'!")
	}

	// Bai 4: Thêm một đối tượng mới vào slice
	data = append(data, Serve{Name: "fileCustome", Class: "org.cofax.cds.FileServlet.Custome"})
	fmt.Println(data)

	// Bai 5: In ra màn hình các địa chỉ của các item strong slice

	for _, i := range data {
		fmt.Println("Name", i, &i.Name)
		fmt.Println("Class", i, &i.Class)
	}

	// Bai 6:
	slide := make([]int, 10)

	mang := [10]int{11, 34, 56, 77, 99, 109, 66, 20, 88, 34}
	for i := 0; i < len(mang)-1; i++ {
		for j := i + 1; j < len(mang); j++ {
			if mang[i] > mang[j] {
				temp := mang[i]
				mang[i] = mang[j]
				mang[j] = temp
			}
		}

	}
	for i := 0; i < len(mang); i++ {
		slide[i] = mang[i]

	}
	fmt.Println(slide)

	cp := slide[1:8]
	fmt.Println(cp)

	// tmp := make([]int, 15)
	// tmp = slide[1:15]

	// fmt.Println(tmp)

	// Bai 7
	bai7()
}
