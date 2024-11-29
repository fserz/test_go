package main

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {
	// byte uint8别名
	// rune int32别名
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	s := "hello seal球仔"
	//// 会乱码
	//for i := 0; i < len(s); i++ {
	//	fmt.Printf("%c\n", s[i])
	//}

	for _, r := range s {
		fmt.Printf("%c\n", r)
	}

	a := 2
	a++
	fmt.Println(a)

	a += 6
	fmt.Println(a)
}
