package main

import (
	"fmt"
)

const (
	minTemp, maxTemp = 15, 30
)

func main() {
	var nDepartments uint

	_, err := fmt.Scan(&nDepartments)
	if err != nil {
		fmt.Println("invalid departments count:", err)

		return
	}

	for range nDepartments {
		fmt.Println("Processing department...")
	}
}
