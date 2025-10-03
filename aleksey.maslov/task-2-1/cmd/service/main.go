package main

import "fmt"

func main() {
	var (
		departmentCount int
		employeesCount  int
	)
	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Invalid department count")
		return
	}
	for range departmentCount {
		_, err = fmt.Scanln(&employeesCount)
		if err != nil {
			fmt.Println("Invalid employees count")
			return
		}
	}
}
