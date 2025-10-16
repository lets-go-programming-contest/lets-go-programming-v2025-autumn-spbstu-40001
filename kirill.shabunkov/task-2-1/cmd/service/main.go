package main

import (
	"fmt"
)

type Temperature struct {
	Min int
	Max int
}

const (
	minTemperature     = 15
	maxTemperature     = 30
	expectedInputCount = 2
)

func (temp *Temperature) getSuitableTemperature(operand string, prefferedTemperature int) {
	switch operand {
	case ">=":
		temp.Min = max(temp.Min, prefferedTemperature)
	case "<=":
		temp.Max = min(temp.Max, prefferedTemperature)
	default:
		fmt.Println("Wrong operand!")

		return
	}

	if temp.Min > temp.Max {
		fmt.Println(-1)
	} else {
		fmt.Println(temp.Min)
	}
}

func main() {
	var departamentNum int

	_, err := fmt.Scan(&departamentNum)
	if err != nil {
		fmt.Println("Invalid input data: ", err)

		return
	}

	for range departamentNum {
		var workerNum int
		currentTemperature := Temperature{Min: minTemperature, Max: maxTemperature}
		_, err := fmt.Scan(&workerNum)
		if err != nil {
			fmt.Println("Invalid input data: ", err)

			return
		}

		for range workerNum {
			var (
				prefferedTemperature int
				operand              string
			)

			count, err := fmt.Scan(&operand, &prefferedTemperature)
			if count != expectedInputCount {
				fmt.Println("Incorrect amount of input data!")

				return
			}

			if err != nil {
				fmt.Println("Invalid input data: ", err)

				return
			}

			currentTemperature.getSuitableTemperature(operand, prefferedTemperature)
		}
	}
}
