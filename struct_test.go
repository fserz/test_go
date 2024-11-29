package main

import (
	"fmt"
	"testing"
)

type person struct {
	name, city string
	age        int8
}

func TestStruct(t *testing.T) {
	var p person
	p.name = "sealball"
	p.city = "sz"
	p.age = 11
	fmt.Printf("p=%#v\n", p)
	fmt.Printf("p=%+v\n", p)

	// 匿名结构体
	var user struct {
		name    string
		married bool
	}
	user.name = "qqq"
	user.married = false
	fmt.Printf("p=%+v\n", user)

	p2 := person{
		name: "ss",
		city: "sz",
		age:  14,
	}
	fmt.Printf("%T\n", p2)

	p3 := &person{
		name: "bb",
		city: "xa",
		age:  20,
	}
	fmt.Printf("%T\n", p3)
}
