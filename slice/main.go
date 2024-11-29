package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
	// s 和 s2 都是基于数组 a 的切片，尽管它们的长度和起始位置不同，但它们共享同一个底层数组。
	// 容量是从切片的起始位置到底层数组末尾的元素个数，而不是切片的长度。
	s3 := s[1:2]
	fmt.Printf("s3:%v len(s3):%v cap(s3):%v\n", s3, len(s3), cap(s3))
}
