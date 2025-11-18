package main

import (
	"fmt"

	"github.com/Nekich06/task-2-1/internal/tmanager"
)

const (
	MaxTemp int = 30
	MinTemp int = 15
)

func main() {
	var deptNum int

	_, err := fmt.Scan(&deptNum)
	if err != nil {
		fmt.Println("cannot scan department number", err)

		return
	}

	for range deptNum {
		var (
			staffNum, wishfulTemp int
			condition             string
		)

		_, err := fmt.Scan(&staffNum)
		if err != nil {
			fmt.Println("cannot scan staff number", err)

			return
		}

		airConditioner := tmanager.New(MaxTemp, MinTemp)

		for range staffNum {
			_, err = fmt.Scan(&condition)
			if err != nil {
				fmt.Println("cannot scan condition", err)

				return
			}

			_, err = fmt.Scan(&wishfulTemp)
			if err != nil {
				fmt.Println("cannot scan wishful temperature", err)

				return
			}

			newOptimalTemp, err := airConditioner.SetAndGetNewOptimalTemp(condition, wishfulTemp)
			if err != nil {
				fmt.Println("failed to recalculate temp", err)

				continue
			}

			fmt.Println(newOptimalTemp)
		}
	}
}
