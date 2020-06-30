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
	//WriteFile()

	file, err := ioutil.ReadFile("serve.json")
	if err != nil {
		fmt.Println(err)
	}

	data := make([]Serve, 2)

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data); i++ {
		log.Println("Name: ", data[i].Name)
		log.Println("Class: ", data[i].Class)
	}

	// Bai 3: lọc kết quả với class có chứa từ `admin` in ra đối tượng phù hợp
	xet := 0
	for i := 0; i < len(data); i++ {

		if strings.Contains(data[i].Class, "admin") {
			fmt.Println("Doi tuong co chu 'admin'la ")
			fmt.Println(data[i])
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

	for i := 0; i < len(data); i++ {
		fmt.Println("Name", i, &data[i].Name)
		fmt.Println("Class", i, &data[i].Class)
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
