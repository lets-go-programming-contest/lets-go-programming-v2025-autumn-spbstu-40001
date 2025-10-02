package main

import (
	"fmt"
	"os"

	"github.com/atroxxxxxx/task-2-1/internal/conditioner"
)

func main() {
	var nDepartments uint

	_, err := fmt.Scan(&nDepartments)
	if err != nil || nDepartments == 0 {
		fmt.Println("invalid departments count")

		return
	}

	for range nDepartments {
		_, err = conditioner.MakeDepartmentTemperature(os.Stdin, os.Stdout)
		if err != nil {
			fmt.Println(err)

			return
		}
	}
}
