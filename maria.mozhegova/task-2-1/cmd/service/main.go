package main

import (
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30
)

type Temperature struct {
	min int
	max int
}

func (temp *Temperature) printTempChange(cond string, currTemp int) {
	switch cond {
	case "<=":
		temp.max = min(temp.max, currTemp)
	case ">=":
		temp.min = max(temp.min, currTemp)
	default:
		fmt.Println("Invalid condition")

		return
	}

	if temp.min > temp.max {
		fmt.Println(-1)
	} else {
		fmt.Println(temp.min)
	}
}

func main() {
	var departNum uint

	_, err := fmt.Scan(&departNum)
	if err != nil {
		fmt.Println("Invalid input:", err)

		return
	}

	for range departNum {
		var (
			employNum   int
			temperature = Temperature{minTemp, maxTemp}
		)
		_, err := fmt.Scan(&employNum)
		if err != nil {
			fmt.Println("Invalid input:", err)

			return
		}

		for range employNum {
			var (
				condition   string
				currentTemp int
			)

			_, err := fmt.Scan(&condition)
			if err != nil {
				fmt.Println("Invalid input:", err)

				return
			}

			_, err = fmt.Scan(&currentTemp)
			if err != nil {
				fmt.Println("Invalid input:", err)

				return
			}

			temperature.printTempChange(condition, currentTemp)
		}
	}
}
