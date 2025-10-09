package main

import "fmt"

func main() {
	var departmentCount int
	_, err := fmt.Scanln(&departmentCount)
	if err != nil {
		fmt.Println("Invalid input", err)
		return
	}
	for range departmentCount {
		var employeeCount int
		_, err := fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("Invalid input", err)
			return
		}
		for range employeeCount {
			var operation string
			var temperature int
			_, err := fmt.Scan(&operation, &temperature)
			if err != nil {
				fmt.Println("Invalid input", err)
				return
			}
		}
	}
}
