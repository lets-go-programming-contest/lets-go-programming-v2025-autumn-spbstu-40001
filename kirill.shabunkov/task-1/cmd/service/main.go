package main

import (
	"errors"
	"fmt"
)

var (
	ErrZeroDivision     = errors.New("Division by zero")
	ErrInvalidOperation = errors.New("Invalid operation")
)

func main() {
	var (
		lhs, rhs  int
		operation string
	)

	scanned, err := fmt.Scan(&lhs, &rhs, &operation)
	switch {
	case err == nil:
	case scanned == 0:
		fmt.Println("Invalid first operand")
		return
	case scanned == 1:
		fmt.Println("Invalid second operand")
		return
	case scanned == 2:
		fmt.Println("Invalid operation")
		return
	}

	res, err := calc(lhs, operation, rhs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func calc(lhs int, operation string, rhs int) (int, error) {
	switch operation {
	case "+":
		return lhs + rhs, nil
	case "-":
		return lhs - rhs, nil
	case "*":
		return lhs * rhs, nil
	case "/":
		if rhs == 0 {
			return 0, ErrZeroDivision
		}
		return lhs / rhs, nil
	}
	return 0, ErrInvalidOperation
}
