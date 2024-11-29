package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//var ch1 chan int
	//ch1 := make(chan int, 1) // 有缓冲区
	ch1 := make(chan int) // 无缓冲区
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-ch1
		fmt.Println(x)
	}()
	ch1 <- 20
	wg.Wait()
	close(ch1)
}
