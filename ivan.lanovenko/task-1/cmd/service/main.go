package main

import "fmt"

func main() {
	var (
		op1, op2  int
		operation string
	)

	if _, err := fmt.Scanln(&op1); err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if _, err := fmt.Scanln(&op2); err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	if _, err := fmt.Scanln(&operation); err != nil {
		fmt.Println("Invalid operation")
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
	}
}
