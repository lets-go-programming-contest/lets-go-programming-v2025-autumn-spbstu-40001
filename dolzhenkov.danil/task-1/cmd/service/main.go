package main

import (
	"fmt"
)

func main() {
	var (
		op1, op2  int
		operation string
	)

	// Читаем первый операнд
	_, err := fmt.Scan(&op1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	// Читаем второй операнд
	_, err = fmt.Scan(&op2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	// Читаем операцию
	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	// Выполняем операцию
	switch operation {
	case "+":
		fmt.Printf("%d\n", op1+op2)
	case "-":
		fmt.Printf("%d\n", op1-op2)
	case "*":
		fmt.Printf("%d\n", op1*op2)
	case "/":
		if op2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Printf("%d\n", op1/op2)
	default:
		fmt.Println("Invalid operation")
	}
}
