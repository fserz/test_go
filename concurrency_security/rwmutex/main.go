package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	lock   sync.Mutex
	rwLock sync.RWMutex
)

func read(wg *sync.WaitGroup) {
	rwLock.RLock()
	rwLock.RLocker()
	rwLock.RLocker()
	defer rwLock.RUnlock()
	//lock.Lock()
	//defer lock.Unlock()
	time.Sleep(time.Millisecond)
	wg.Done()
}

func write(wg *sync.WaitGroup) {
	lock.Lock()
	defer lock.Unlock()
	x++
	time.Sleep(time.Millisecond * 10)
	wg.Done()
}

func main() {
	//调用一次
	//var once syncMap.Once
	//once.Do()
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read(&wg)
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write(&wg)
	}
	wg.Wait()
	//duration := time.Since(start)
	duration := time.Now().Sub(start)
	fmt.Println(duration)
}
