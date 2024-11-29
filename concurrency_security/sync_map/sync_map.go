package main

import (
	"fmt"
	"sync"
)

// sync.Map 并发安全的Map
var m2 = sync.Map{}

// 使用普通map并发写会panic
var m = make(map[int]int)

var wg sync.WaitGroup

func get(key int) int {
	return m[key]
}

func set(key, value int) {
	m[key] = value
}

// 使用普通map并发写会panic
//func main() {
//	for i := 0; i < 20; i++ {
//		wg.Add(1)
//		go func(i int) {
//			set(i, i+200)
//			fmt.Printf("key:%d value:%v", i, get(i))
//		}(i)
//	}
//	wg.Wait()
//}

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			m2.Store(i, i+200)
			value, _ := m2.Load(i)
			fmt.Printf("key:%d value:%v\n", i, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
