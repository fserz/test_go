package _if

import (
	"fmt"
	"testing"
)

func TestIf(t *testing.T) {
	if score := 12; score >= 23 {
		fmt.Println("A")
	} else if score <= 10 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
