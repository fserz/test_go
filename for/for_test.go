package _for

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	var i = 0
	for ; i < 10; i++ {
		switch i {
		case 1, 3, 5, 7, 9:
			fmt.Printf("%d 奇数 \n", i)
		case 2, 4, 6, 8:
			fmt.Printf("%d 偶数 \n", i)
		}
	}
	age := 29
	switch {
	case age > 18:
		fmt.Println("warning")
	default:
		fmt.Println("嘿嘿")
	}

	var boolArray = [...]bool{true, false, true}
	var langArray = [...]string{0: "Golang", 1: "Python", 2: "Java"}
	fmt.Println(len(boolArray))
	fmt.Println(boolArray)
	fmt.Println(len(langArray))
	fmt.Println(langArray)

	for i, value := range boolArray {
		fmt.Println(i, " + ", value)
	}

	cityArray := [3][2]string{
		{"beijing", "shanghai"},
		{"chongqing", "xian"},
		{"test", "shenzhen"},
	}

	for i, value := range cityArray {
		fmt.Println(i, "+", value)
	}

	for _, v1 := range cityArray {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//fmt.Println(cityArray)
	//fmt.Println(cityArray[1][1])

}
