package main

import (
	"errors"
	"fmt"
)

func calculator(a float64, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return (a + b), nil
	case "-":
		return (a - b), nil
	case "*":
		return (a * b), nil
	case "/":
		if b != 0 {
			return (a / b), nil
		}
		return 0, errors.New("division by zero")
	default:
		return 0, errors.New("invalid operation")
	}
}

func main() {
	var a float64 = 0.0
	var b float64 = 0.0
	var op string
	_, err1 := fmt.Scan(&a)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err2 := fmt.Scan(&b)
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}
	fmt.Scan(&op)
	res, err := calculator(a, b, op)
	if err == nil {
		fmt.Println(res)
		return
	}
	fmt.Println(err)
}
