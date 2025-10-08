package main

import (
	"fmt"
	"io"

	"github.com/Aapng-cmd/task-2-1/internal/freezer"
)

func main() {
	var (
		nNumberOfOlimpic int
		reader           io.Reader
	)

	_, err := fmt.Scan(&nNumberOfOlimpic)
	if err != nil || nNumberOfOlimpic < 0 {
		fmt.Println("Invalid departure count")

		return
	}

	err = freezer.CalcForDepartment(reader, nNumberOfOlimpic)
	if err != nil {
		fmt.Println(err)
	}
}
