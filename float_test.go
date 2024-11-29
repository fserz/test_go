package main

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestFloat(t *testing.T) {
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	str := "hello seal ball ball"
	s := strings.Split(str, " ")
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[3])
	fmt.Println(strings.Contains(str, "seal"))
	fmt.Println(strings.HasPrefix(str, "hello "))
	fmt.Println(strings.Index(str, "sea"))
	s4 := []string{"hello", "cute", "seal"}
	fmt.Println(strings.Join(s, "---"))
	fmt.Println(strings.Join(s4, "+++"))
}
