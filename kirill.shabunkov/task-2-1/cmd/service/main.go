package main

import (
	"fmt"
)

type Temperature struct {
	Min int
	Max int
}

func NewTemperature() Temperature {
	return Temperature{
		Min: minTemperature,
		Max: maxTemperature,
	}
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
		var (
			currentTemperature = NewTemperature()
			workerNum          int
		)

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
			if err != nil {
				fmt.Printf("Error reading input: expected %d values, got error: %v\n", expectedInputCount, err)

				return
			}

			if count != expectedInputCount {
				fmt.Printf("Incorrect amount of input data: expected %d values, but got %d\n", expectedInputCount, count)

				return
			}

			currentTemperature.getSuitableTemperature(operand, prefferedTemperature)
		}
	}
}
