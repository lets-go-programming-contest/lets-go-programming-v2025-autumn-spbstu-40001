package main

import (
	"fmt"
	"os"

	"github.com/Aapng-cmd/task-2-1/internal/freezer"
)

func main() {
	var nNumberOfOlimpic int

	_, err := fmt.Scan(&nNumberOfOlimpic)
	if err != nil || nNumberOfOlimpic < 0 {
		fmt.Println("Invalid departure count")

		return
	}

	err = freezer.CalcForDepartment(os.Stdin, nNumberOfOlimpic)
	if err != nil {
		fmt.Println(err)
	}
}
