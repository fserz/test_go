package main

import "sync"

var (
	x    int64
	lock sync.Mutex
)

func add(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 70000; i++ {
		lock.Lock() // 加锁
		x++
		lock.Unlock() // 解锁
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go add(&wg)
	go add(&wg)
	wg.Wait()
	println(x)
}
