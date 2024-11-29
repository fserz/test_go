package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	var a = [...]int{7, 12, 1, 9, 5}
	// a[:]得到的是一个切片，指向底层数组a
	sort.Ints(a[:])
	fmt.Println(a)
}
