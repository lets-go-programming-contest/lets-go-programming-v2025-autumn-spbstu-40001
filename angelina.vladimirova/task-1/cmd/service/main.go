package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Enter first operand: ")
	aStr, _ := r.ReadString('\n')
	aStr = strings.TrimSpace(aStr)

	a, err := strconv.Atoi(aStr)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	fmt.Print("Enter second operand: ")
	bStr, _ := r.ReadString('\n')
	bStr = strings.TrimSpace(bStr)

	b, err := strconv.Atoi(bStr)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	fmt.Print("Enter operation (+, -, *, /): ")
	op, _ := r.ReadString('\n')
	op = strings.TrimSpace(op)

	var res int

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
			return
		}
		res = a / b
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(res)
}
