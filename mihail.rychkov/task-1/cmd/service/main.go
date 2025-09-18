package main;

import "fmt";

func evaluate(lhs int, operation string, rhs int) (int, error) {
	switch operation {
	case "+":
		return lhs + rhs, nil;
	case "-":
		return lhs - rhs, nil;
	case "*":
		return lhs * rhs, nil;
	case "/":
		if rhs == 0 {
			return 0, fmt.Errorf("Division by zero");
		}
		return lhs / rhs, nil;
	}
	return 0, fmt.Errorf("Invalid operation");
}

func main() {
	var (
		lhs, rhs int;
		operation string;
	);

	scanned, _ := fmt.Scan(&lhs, &rhs, &operation);
	switch scanned {
	case 0:
		fmt.Println("Invalid first operand");
		return;
	case 1:
		fmt.Println("Invalid second operand");
		return;
	case 2:
		fmt.Println("Invalid operation");
		return;
	}

	res, err := evaluate(lhs, operation, rhs);
	if err != nil {
		fmt.Println(err);
		return;
	}
	fmt.Println(res);
}
