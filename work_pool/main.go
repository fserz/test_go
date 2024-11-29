package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d start job: %d\n", id, job)
		results <- job * 2
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("worker:%d stop job: %d\n", id, job)
	}
}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 3个goroutinue
	for i := 0; i < 3; i++ {
		wg.Add(1) // 为每个 worker 增加计数
		go worker(i, jobs, results)
	}
	// 5个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	for result := range results {
		fmt.Println(result)
	}

}
