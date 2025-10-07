package main

import (
	"fmt"
	"os"

	"github.com/atroxxxxxx/task-2-1/internal/conditioner"
)

func main() {
	var nDepartments uint

	_, err := fmt.Scan(&nDepartments)
	if err != nil {
		fmt.Println("invalid departments count:", err)

		return
	}

	for range nDepartments {
		_, err = conditioner.CalcDepartmentTemperature(os.Stdin, os.Stdout)
		if err != nil {
			fmt.Println("failed calculate department temperature:", err)

			return
		}
	}
}
