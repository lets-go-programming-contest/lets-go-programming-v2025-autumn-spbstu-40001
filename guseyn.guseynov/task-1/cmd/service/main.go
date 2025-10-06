package main

import (
	"errors"
	"fmt"
)

// Calculate выполняет математическую операцию и возвращает результат или ошибку
func Calculate(operation string, op1, op2 int) (int, error) {
	switch operation {
	case "+":
		return op1 + op2, nil
	case "*":
		return op1 * op2, nil
	case "-":
		return op1 - op2, nil
	case "/":
		if op2 == 0 {
			return 0, errors.New("division by zero")
		}
		return op1 / op2, nil
	default:
		return 0, errors.New("invalid operation")
	}
}

func main() {
	var (
		op1, op2  int
		operation string
	)

	_, err := fmt.Scan(&op1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&op2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	result, err := Calculate(operation, op1, op2)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Println(result)
}
