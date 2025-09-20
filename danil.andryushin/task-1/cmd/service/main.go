package main

import (
	"errors"
	"fmt"
)

var (
	ErrDivByZero = errors.New("Division by zero")
	ErrInvalidOp = errors.New("Invalid operation")
)

func calc(a, b int, operation string) (int, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, ErrDivByZero
		}
		return a / b, nil
	default:
		return 0, ErrInvalidOp
	}
}

func main() {
	var (
		lhs, rhs  int
		operation string
	)
	scanned, err := fmt.Scan(&lhs, &rhs, &operation)
	if err != nil {
		switch scanned {
		case 0:
			fmt.Println("Invalid first operand")
		case 1:
			fmt.Println("Invalid second operand")
		case 2:
			fmt.Println("Operation input error")
		}
	}
	result, err := calc(lhs, rhs, operation)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
