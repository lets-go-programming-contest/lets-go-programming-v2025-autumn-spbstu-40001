package main

import "fmt"

func main() {
	var departmentCount int
	_, err := fmt.Scanln(&departmentCount)
	if err != nil {
		fmt.Println("Invalid input", err)
		return
	}
	fmt.Printf("Прочитано отделов: %d\n", departmentCount)
}
