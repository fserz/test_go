package main

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	a1()
	b()
	c()
}

func a1() {
	fmt.Println("func a")
}

func b() {
	// recover 要在可能触发panic的语句之前
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func b err")
		}
	}()
	panic("panic b")
}

func c() {
	fmt.Println("func c")
}
