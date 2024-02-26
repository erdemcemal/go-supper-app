package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func Multiply(a, b int) int {
	return a * b
}

func Subtract(a, b int) int {
	return a - b
}

func main() {
	fmt.Println(Add(1, 2))

	res, _ := Divide(10, 2)

	fmt.Println(res)
}
