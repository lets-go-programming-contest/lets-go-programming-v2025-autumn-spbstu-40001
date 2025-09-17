package main

import (
	"fmt"
)

func main() {
	var firstOp int
	var secondOp int
	var operation string

	_, err := fmt.Scan(&firstOp)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&secondOp)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	fmt.Scan(&operation)

	var result int

	switch operation {
	case "+":
		result = firstOp + secondOp
	case "-":
		result = firstOp - secondOp
	case "*":
		result = firstOp * secondOp
	case "/":
		if secondOp != 0 {
			result = firstOp / secondOp
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
