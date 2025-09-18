package main

import "fmt"

func calc(a, b int, operation string) (int, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Invalid operation")
	}
}

func main() {
	var (
		lhs, rhs  int    = 0, 0
		operation string = ""
	)
	scanned, _ := fmt.Scan(&lhs, &rhs, &operation)
	switch scanned {
	case 0:
		fmt.Println("Invalid first operand")
	case 1:
		fmt.Println("Invalid second operand")
	case 2:
		fmt.Println("Operation input error")
	default:
		result, err := calc(lhs, rhs, operation)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	}
}
