package main

import "fmt"

// 1. 生成0~100发送到ch1
// 2. 从ch1中取出数据计算他的平方，结果发送到ch2

// 1. 生成0~100发送到ch1
// chan-< 表示为单向通道，只能往里面写
// <-chan 表示只能往外读的单向通道
func f1(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

// 2. 从ch1中取出数据计算他的平方，结果发送到ch2
func f2(ch1 <-chan int, ch2 chan<- int) {
	defer close(ch2)
	// 从通道取值两种方式
	//for {
	//	temp, ok := <-ch1
	//	if !ok {
	//		break
	//	}
	//	ch2 <- temp * temp
	//}
	// 第二种
	for temp := range ch1 {
		ch2 <- temp * temp
	}

}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go f1(ch1)
	go f2(ch1, ch2)
	// 第二种chan种取值方式
	for ret := range ch2 {
		fmt.Println(ret)
	}
}
