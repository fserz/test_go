package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	id   int
	name string
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{
			id:   666,
			name: "sealzeng",
		}
	})
	return instance
}

func main() {
	ins := GetInstance()
	fmt.Printf("%+v", ins)
}
