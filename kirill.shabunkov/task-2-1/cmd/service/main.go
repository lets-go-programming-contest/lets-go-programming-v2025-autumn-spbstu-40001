package main

import (
	"errors"
	"fmt"
)

const (
	minTemperature = 15
	maxTemperature = 30
)

var ErrInvalidOperand = errors.New("invalid operand")

type Temperature struct {
	Min int
	Max int
}

func NewTemperature(minTemperature int, maxTemperature int) Temperature {
	return Temperature{
		Min: minTemperature,
		Max: maxTemperature,
	}
}

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
			currentTemperature = NewTemperature(minTemperature, maxTemperature)
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

			_, err := fmt.Scan(&operand, &prefferedTemperature)
			if err != nil {
				fmt.Printf("Input error: %v\n", err)

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
