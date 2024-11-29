package main

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	res := intSum(12, 14, 53) // a = 12, b = [14, 53]
	fmt.Println(res)

	abc := testGlobal
	fmt.Printf("%T\n", abc)
	abc()

	r1 := calc(10, 20, add)
	r2 := calc(10, 20, sub)

	fmt.Println(r1)
	fmt.Println(r2)

}

func intSum(a int, b ...int) int {
	res := a
	for _, arg := range b {
		res += arg
	}
	return res
}

func testGlobal() {
	num := 100
	name := "seal"
	fmt.Println("变量num", num)
	fmt.Println(name)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func BenchmarkGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testGlobal()
	}
}
