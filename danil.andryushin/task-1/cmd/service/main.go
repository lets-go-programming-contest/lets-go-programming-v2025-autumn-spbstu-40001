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

	_, err := fmt.Scan(&lhs)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&rhs)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	result, err := calc(lhs, rhs, operation)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
