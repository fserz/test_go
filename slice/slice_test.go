package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// 基于数组得到切片
	a := [5]int{1, 2, 3, 4, 5}
	b := a[1:4]
	fmt.Println(b)
	fmt.Printf("a - %T\n", a)
	fmt.Printf("b - %T\n", b)
	// make构造切片
	d := make([]int, 5, 10)
	fmt.Printf("d - %T\n", d)
	// len(d) 获取容量
	fmt.Println(len(d))
	// cap(d)获取容量
	fmt.Println(cap(d))

	var c []int     // 声明int类型切片
	var e = []int{} // 声明并初始化
	if c == nil {
		fmt.Println("c是nil")
	}
	if e == nil {
		fmt.Println("e是nil")
	}
	fmt.Println(c, len(c), cap(c))
	fmt.Println(e, len(e), cap(e))

	var f []int
	f = append(f, 12)
	fmt.Println(f[0])
	fmt.Println(f)
	for i := 0; i < 10; i++ {
		f = append(f, i, i+1)
		fmt.Printf("%v len:%d cap: %d ptr:%p\n", f, len(f), cap(f), f)
	}

	g := []int{1, 2, 3, 4, 5}
	h := []int{6, 7, 8, 9}
	g = append(g, h...)
	fmt.Println(g)

	// 切片修改相同指向的会影响
	fmt.Println(h)
	i := h
	fmt.Println(i)
	i[0] = 99
	fmt.Println(h)
	fmt.Println(i)

	j := []int{1, 2, 3, 4, 5}
	k := make([]int, 5, 5)
	l := k
	// copy拷贝了一个新的
	copy(k, j)
	k[0] = 100
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	// 切片删除元素
	m := []string{"bj", "sh", "gz", "sz"}
	m = append(m[0:2], m[3:]...)
	fmt.Println(m)
}
