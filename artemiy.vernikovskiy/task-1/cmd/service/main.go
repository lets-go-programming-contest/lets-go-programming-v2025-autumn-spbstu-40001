package main

import (
    "fmt"
)

func main() {
    var a int = 0
    var b int = 0
    
    _, err := fmt.Scan(&a)
    if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
    _, err := fmt.Scan(&b)
    if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	
	var oper string
	_, err := fmt.Scan(&oper)
    if err != nil {
		fmt.Println("Invalid operation")
		return
	}
	
	switch oper {
	case "+":
	    fmt.Println(a + b)
	case "-":
	    fmt.Println(a - b)
	case "*":
	    fmt.Println(a + b)
	case "/":
	    switch b {
	    case 0:
	        fmt.Println("Division by zero")
	    default:
	        fmt.Println(a / b)
	    }
	default:
	    fmt.Println("Invalid operation")
	}
}
