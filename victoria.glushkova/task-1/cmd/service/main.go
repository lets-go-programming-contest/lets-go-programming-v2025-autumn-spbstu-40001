package main

import (
	"fmt"
)

func readInt() (int, bool) {
	var value int
	if _, err := fmt.Scan(&value); err != nil {
		return 0, false
	}
	return value, true
}

func readOperation() (string, bool) {
	var op string
	if _, err := fmt.Scan(&op); err != nil {
		return "", false
	}
	return op, true
}

func calculate(a, b int, operation string) string {
	switch operation {
	case "+":
		return fmt.Sprintf("%d", a+b)
	case "-":
		return fmt.Sprintf("%d", a-b)
	case "*":
		return fmt.Sprintf("%d", a*b)
	case "/":
		if b == 0 {
			return "Division by zero"
		}
		return fmt.Sprintf("%d", a/b)
	default:
		return "Invalid operation"
	}
}

func main() {
	first, ok := readInt()
	if !ok {
		fmt.Println("Invalid first operand")
		return
	}

	second, ok := readInt()
	if !ok {
		fmt.Println("Invalid second operand")
		return
	}

	operation, ok := readOperation()
	if !ok {
		fmt.Println("Invalid operation")
		return
	}

	result := calculate(first, second, operation)
	fmt.Println(result)
}
