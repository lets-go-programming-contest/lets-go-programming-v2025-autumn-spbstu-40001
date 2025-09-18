package main;

import "fmt";

func main() {
	var (
		lhs, rhs int = 0, 0;
		operation string = "";
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
}
