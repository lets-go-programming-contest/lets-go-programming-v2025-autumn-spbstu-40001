package main

import "fmt"

func main() {
 var firstOperand, secondOperand int
 var operation string

 fmt.Scan(&firstOperand)
 fmt.Scan(&secondOperand)
 fmt.Scan(&operation)

 if operation == "/" && secondOperand == 0 {
  fmt.Println("Division by zero")
  return
 }

 switch operation {
 case "+":
  fmt.Println(firstOperand + secondOperand)
 case "-":
  fmt.Println(firstOperand - secondOperand)
 case "*":
  fmt.Println(firstOperand * secondOperand)
 case "/":
  fmt.Println(firstOperand / secondOperand)
 default:
  fmt.Println("Invalid operation")
 }
}
