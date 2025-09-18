package calculator

import "errors"

func Sum(a int, b int) int {
	return a + b
}

func Dif(a int, b int) int {
	return a - b
}

func Div(a int, b int) (int, error) {
	if b == 0 {
		return 1, errors.New("Division by zero")
	}
	return a / b, nil
}

func Mul(a int, b int) int {
	return a * b
}

func isOperator(op rune) bool {
	return (op == '+' || op == '-' || op == '*' || op == '/')
}

func IsOperator(op rune) bool {
	return isOperator(op)
}
