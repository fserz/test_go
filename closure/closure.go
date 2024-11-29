package main

import "fmt"

// 捕获外部变量
//func main() {
//	// 定义一个外部变量
//	x := 10
//
//	// 定义一个闭包，引用外部变量x
//	closure := func(y int) int {
//		return x + y
//	}
//
//	fmt.Println(closure(5)) // 输出：15
//
//	// 修改外部变量x的值
//	x = 20
//	fmt.Println(closure(5)) // 输出：25
//}

// 创建一个闭包，每次调用返回一个累加器函数
func createAdder(base int) func(int) int {
	return func(delta int) int {
		base += delta // 修改捕获的外部变量base
		return base
	}
}

func main() {
	adder1 := createAdder(10) // 创建一个累加器
	fmt.Println(adder1(5))    // 输出：15
	fmt.Println(adder1(10))   // 输出：25

	adder2 := createAdder(100) // 创建另一个累加器
	fmt.Println(adder2(10))    // 输出：110
	fmt.Println(adder1(5))     // 输出：30 （adder1与adder2互不影响）
}
