package main

import (
	"fmt"
)

func main() {
	runtime.version()

	var first_op int
	var second_op int
	var operation string

	_, err := fmt.Scan(&first_op)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&second_op)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	fmt.Scan(&operation)

	var result int

	switch operation {
	case "+":
		result = first_op + second_op
	case "-":
		result = first_op - second_op
	case "*":
		result = first_op * second_op
	case "/":
		if second_op != 0 {
			result = first_op / second_op
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
