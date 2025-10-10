package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidOp    = errors.New("Invalid operation")
	ErrDivideByZero = errors.New("Division by zero")
)

func evaluate(operand1, opearand2 int, op string) (int, error) {
	switch op {
	case "+":
		return operand1 + opearand2, nil
	case "-":
		return operand1 - opearand2, nil
	case "*":
		return operand1 * opearand2, nil
	case "/":
		if opearand2 == 0 {
			return 0, ErrDivideByZero
		}
		return operand1 / opearand2, nil
	}
	return 0, ErrInvalidOp
}

func main() {
	var (
		operand1, operand2 int
		op                 string
	)

	_, err := fmt.Scan(&operand1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&operand2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, err = fmt.Scan(&op)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	result, err := evaluate(operand1, operand2, op)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
