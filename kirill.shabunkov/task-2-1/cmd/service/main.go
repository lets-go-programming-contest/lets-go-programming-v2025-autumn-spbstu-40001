package main

import (
	"errors"
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

var ErrInvalidOperand = errors.New("invalid operand")

func (temp *Temperature) getSuitableTemperature(operand string, prefferedTemperature int) (int, error) {
	switch operand {
	case ">=":
		temp.Min = max(temp.Min, prefferedTemperature)
	case "<=":
		temp.Max = min(temp.Max, prefferedTemperature)
	default:
		return 0, fmt.Errorf("%w: %s, expected '>=' or '<='", ErrInvalidOperand, operand)
	}

	if temp.Min > temp.Max {
		return -1, nil
	}

	return temp.Min, nil
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

			result, err := currentTemperature.getSuitableTemperature(operand, prefferedTemperature)
			if err != nil {
				fmt.Println("Error: ", err)

				return
			}

			fmt.Println(result)
		}
	}
}
