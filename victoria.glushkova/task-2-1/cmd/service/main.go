package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	minTemperature = 15
	maxTemperature = 30
)

var ErrInvalidOperation = errors.New("invalid operation")

type Temperature struct {
	Min int
	Max int
}

func NewTemperature(minTemperature, maxTemperature int) Temperature {
	return Temperature{
		Min: minTemperature,
		Max: maxTemperature,
	}
}

func (temp *Temperature) getSuitableTemperature(operand string, preferredTemperature int) (int, error) {
	if temp.Min > temp.Max {
		return -1, nil
	}

	switch operand {
	case ">=":
		if preferredTemperature > temp.Max {
			temp.Min = 1
			temp.Max = 0

			return -1, nil
		}

		if preferredTemperature > temp.Min {
			temp.Min = preferredTemperature
		}
	case "<=":
		if preferredTemperature < temp.Min {
			temp.Min = 1
			temp.Max = 0

			return -1, nil
		}

		if preferredTemperature < temp.Max {
			temp.Max = preferredTemperature
		}
	default:
		return 0, fmt.Errorf("%w: %s", ErrInvalidOperation, operand)
	}

	if temp.Min > temp.Max {
		return -1, nil
	}

	return temp.Min, nil
}

func main() {
	var departmentNum int

	_, err := fmt.Scan(&departmentNum)
	if err != nil {
		os.Exit(1)
	}

	for range departmentNum {
		var workerNum int

		_, err := fmt.Scan(&workerNum)
		if err != nil {
			os.Exit(1)
		}

		currentTemperature := NewTemperature(minTemperature, maxTemperature)

		for range workerNum {
			var (
				preferredTemperature int
				operand              string
			)

			_, err := fmt.Scan(&operand, &preferredTemperature)
			if err != nil {
				os.Exit(1)
			}

			result, err := currentTemperature.getSuitableTemperature(operand, preferredTemperature)
			if err != nil {
				os.Exit(1)
			}

			fmt.Println(result)
		}
	}
}
