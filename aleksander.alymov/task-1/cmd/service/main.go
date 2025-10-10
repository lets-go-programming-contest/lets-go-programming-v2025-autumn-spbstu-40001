package main

import "fmt"

func main() {
	var (
		input1, input2 int
		operation      string
	)

	_, err := fmt.Scan(&input1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&input2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Operation input error")
		return
	}

	switch operation {
	case "+":
		fmt.Println(input1 + input2)
	case "-":
		fmt.Println(input1 - input2)
	case "*":
		fmt.Println(input1 * input2)
	case "/":
		if input2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(input1 / input2)
	default:
		fmt.Println("Invalid operation")
	}
}
