package main

import (
	"fmt"
)

func main() {
	var op1, op2 float32
	var oper string
	_, err1 := fmt.Scan(&op1)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}
	fmt.Scan(&oper)
	_, err2 := fmt.Scan(&op2)
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var result float32
	switch oper {
	case "+":
		result = op1 + op2
	case "-":
		result = op1 - op2
	case "*":
		result = op1 * op2
	case "/":
		if op2 != 0 {
			result = op1 / op2
		} else {
			fmt.Println("Division by zero")
			return
		}
	default:
		fmt.Println("Invalid operation")
		return
	}
	fmt.Println(result)
}
