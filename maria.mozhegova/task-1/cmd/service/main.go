package main

import (
	"fmt"
)

func main() {
	var (
		op1, op2  int
		operation string
	)
	_, err1 := fmt.Scan(&op1)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err2 := fmt.Scan(&op2)
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, err3 := fmt.Scan(&operation)
	if err3 != nil {
		fmt.Println("Invalid input operation")
		return
	}

	switch operation {
	case "+":
		fmt.Println(op1 + op2)
	case "-":
		fmt.Println(op1 - op2)
	case "*":
		fmt.Println(op1 * op2)
	case "/":
		if op2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(op1 / op2)
	default:
		fmt.Println("Invalid operation")
		return
	}
}
