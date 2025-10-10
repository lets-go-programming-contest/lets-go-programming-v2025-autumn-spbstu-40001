package main

import (
	"fmt"

	"github.com/Nekich06/task-2-1/internal/temp_manager"
)

const (
	MaxTemp int = 30
	MinTemp int = 15
)

func main() {
	var deptNum int

	_, err := fmt.Scan(&deptNum)
	if err != nil {
		fmt.Println("scan department number error")

		return
	}

	for range deptNum {
		var (
			staffNum, wishfulTemp int
			condition             string
		)

		_, err := fmt.Scan(&staffNum)
		if err != nil {
			fmt.Println("scan staff number error")

			return
		}

		var airConditioner temp_manager.TempManager

		airConditioner.Init(MaxTemp, MinTemp)

		for range staffNum {
			_, err = fmt.Scan(&condition)
			if err != nil {
				fmt.Println("scan condition error")

				return
			}

			_, err = fmt.Scan(&wishfulTemp)
			if err != nil {
				fmt.Println("scan wishful temp error")

				return
			}

			err := airConditioner.SetNewOptimalTemp(condition, wishfulTemp)
			if err != nil {
				fmt.Println(-1)
			} else {
				fmt.Println(airConditioner.GetCurrentOptimalTemp())
			}
		}
	}
}
