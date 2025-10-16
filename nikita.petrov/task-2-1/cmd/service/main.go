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
		fmt.Println(err)

		return
	}

	for range deptNum {
		var (
			staffNum, wishfulTemp int
			condition             string
		)

		_, err := fmt.Scan(&staffNum)
		if err != nil {
			fmt.Println(err)

			return
		}

		airConditioner := tmanager.New(MaxTemp, MinTemp)

		for range staffNum {
			_, err = fmt.Scan(&condition)
			if err != nil {
				fmt.Println(err)

				return
			}

			_, err = fmt.Scan(&wishfulTemp)
			if err != nil {
				fmt.Println(err)

				return
			}

			fmt.Println(airConditioner.SetAndGetNewOptimalTemp(condition, wishfulTemp))
		}
	}
}
