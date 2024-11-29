package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Hello(i int) {
	fmt.Println("hello ball", i)
	wg.Done()
}

func main() {
	//wg.Add(10000)
	//for i := 0; i < 10000; i++ {
	//	wg.Add(1)
	//	go func() {
	//		fmt.Println("hello ball", i)
	//		wg.Done()
	//	}()
	//}

	//for i := 0; i < 10000; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		fmt.Println("hello ball", i)
	//		wg.Done()
	//	}(i)
	//}
	//fmt.Println("hello main")
	//wg.Wait()

	wg.Add(2)
	//runtime.GOMAXPROCS(2)
	go a11()
	go b21()
	wg.Wait()
}

func a11() {
	for i := 0; i < 15; i++ {
		fmt.Println("A: ", i)
	}
	wg.Done()
}

func b21() {
	for i := 0; i < 15; i++ {
		fmt.Println("B: ", i)
	}
	wg.Done()
}
