package main

import (
	"errors"
	"fmt"
)

var (
	ErrZeroDivision     = errors.New("Division by zero")
	ErrInvalidOperation = errors.New("Invalid operation")
)

func calculate(op1 int, operation string, op2 int) (int, error) {
	switch operation {
	case "+":
		return op1 + op2, nil
	case "-":
		return op1 - op2, nil
	case "*":
		return op1 * op2, nil
	case "/":
		if op2 == 0 {
			return 0, ErrZeroDivision
		}
		return op1 / op2, nil
	}
	return 0, ErrInvalidOperation
}

func main() {
	var (
		op1, op2  int
		operation string
	)

	scanned, err := fmt.Scan(&op1, &op2, &operation)
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

	res, err := calculate(op1, operation, op2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
