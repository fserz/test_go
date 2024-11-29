package main

import (
	"fmt"
	"strings"
	"testing"
)

// 匿名函数和闭包
func TestFunc1(t *testing.T) {
	func() {
		fmt.Println("test seal ...")
	}()
	r := a()
	r() // 相当于执行了a函数内部的匿名函数

	// 闭包 = 函数 + 外层变量的引用
	z := makeSuffixFunc(".txt")
	ret := z("sealball")
	fmt.Println(ret)
}

func a() func() {
	return func() {
		fmt.Println("hello seal...")
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
