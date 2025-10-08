package main

import (
	"fmt"
)

func main() {
	var (
		a, b int
		op   string
	)

	_, err := fmt.Scan(&a, &b, &op)
	if err != nil {
		fmt.Println("Input error")
		return
	}

	result := 0
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = a / b
	default:
		fmt.Println("Unknown operation")
		return
	}

	fmt.Print(result)
}
