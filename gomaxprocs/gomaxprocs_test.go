package gomaxprocs

import (
	"fmt"
	"runtime"
	"testing"
)

func a121() {
	for i := 0; i < 15; i++ {
		fmt.Println("A: ", i)
	}
}

func b121() {
	for i := 0; i < 15; i++ {
		fmt.Println("B: ", i)
	}
}

func TestGomaxprocs(t *testing.T) {
	runtime.GOMAXPROCS(1)
	go a121()
	go b121()
}
