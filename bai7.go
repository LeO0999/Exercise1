package main

import (
	"fmt"
	"strconv"
)

type User struct {
	name    string
	age     int64
	gender  bool
	address string
}

func (u *User) GetName() string {
	return u.name
}
func (u *User) GetAge() int64 {
	return u.age
}
func (u *User) GetGender() bool {
	return u.gender
}
func (u *User) GetAddress() string {
	return u.address
}

func bai7() {

	u1 := User{name: "K1", age: 20, gender: true, address: "Paris"}
	u2 := User{name: "K2", age: 25, gender: true, address: "NYC"}
	u3 := User{name: "K3", age: 20, gender: false, address: "Paris"}
	u4 := User{name: "K4", age: 27, gender: true, address: "London"}
	u5 := User{name: "K5", age: 40, gender: true, address: "HCM"}
	u6 := User{name: "K6", age: 60, gender: true, address: "HN"}
	u7 := User{name: "K7", age: 20, gender: false, address: "Paris"}
	u8 := User{name: "K8", age: 27, gender: true, address: "London"}
	u9 := User{name: "K9", age: 40, gender: true, address: "HCM"}
	u10 := User{name: "K10", age: 60, gender: true, address: "HN"}
	m := make(map[string]User)

	m["K1"] = u1
	m["K2"] = u2
	m["K3"] = u3
	m["K4"] = u4
	m["K5"] = u5
	m["K6"] = u6
	m["K7"] = u7
	m["K8"] = u8
	m["K9"] = u9
	m["K10"] = u10

	s := make([]User, 10)
	for i := 0; i < len(s); i++ {
		t := "K" + strconv.Itoa(i+1)
		s[i] = m[t]
	}

	for _, i := range s {
		fmt.Println(i)
	}
}
