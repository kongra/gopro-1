package routes

import (
	"fmt"
)

const (
	SomeValue = 627
)

func Fibonacci(n uint16) (uint64, error) {
	if n > 93 {
		return 0, fmt.Errorf("value out of range %d", n)
	}
	var a, b uint64
	a = 0
	b = 1
	for range n {
		a, b = b, a+b
	}
	return a, nil
}
