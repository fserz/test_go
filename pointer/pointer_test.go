package pointer

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	a := 10
	b := &a
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", *b)
	fmt.Printf("%d\n", *b)

	fmt.Println("----------------------")
	a = 1
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
}

func modify1(x int) {
	x = 100
}

func modify2(y *int) {
	*y = 100
}
