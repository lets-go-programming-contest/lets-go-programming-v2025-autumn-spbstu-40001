package main

import (
	"errors"
	"fmt"
)

var (
	ErrDivisionByZero = errors.New("Division by zero")
	ErrInvalidOperation = errors.New("Invalid operation")
)

func compute(firstOperand int, operator string, secondOperand int) (int, error) {
	switch operator {
	case "+":
		return firstOperand + secondOperand, nil
	case "-":
		return firstOperand - secondOperand, nil
	case "*":
		return firstOperand * secondOperand, nil
	case "/":
		if secondOperand == 0 {
			return 0, ErrDivisionByZero
		}
		return firstOperand / secondOperand, nil
	}
	return 0, ErrInvalidOperation
}

func main() {
	var (
		firstNum, secondNum int
		operator string
	)

	scannedCount, scanErr := fmt.Scan(&firstNum, &secondNum, &operator)
	switch {
	case scanErr == nil:
	case scannedCount == 0:
		fmt.Println("Invalid first operand")
		return
	case scannedCount == 1:
		fmt.Println("Invalid second operand")
		return
	case scannedCount == 2:
		fmt.Println("Invalid operation")
		return
	}

	result, calcErr := compute(firstNum, operator, secondNum)
	if calcErr != nil {
		fmt.Println(calcErr)
		return
	}
	fmt.Println(result)
}
