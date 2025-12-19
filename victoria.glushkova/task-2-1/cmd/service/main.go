package main

import (
	"fmt"
)

const (
	minTemperature = 15
	maxTemperature = 30
)

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

func (temp *Temperature) getSuitableTemperature(operand string, preferredTemperature int) int {
	if temp.Min > temp.Max {
		return -1
	}

	switch operand {
	case ">=":
		if preferredTemperature > temp.Max {
			temp.Min = 1
			temp.Max = 0
			return -1
		}

		if preferredTemperature > temp.Min {
			temp.Min = preferredTemperature
		}
	case "<=":
		if preferredTemperature < temp.Min {
			temp.Min = 1
			temp.Max = 0
			return -1
		}

		if preferredTemperature < temp.Max {
			temp.Max = preferredTemperature
		}
	default:
		temp.Min = 1
		temp.Max = 0
		return -1
	}

	if temp.Min > temp.Max {
		return -1
	}

	return temp.Min
}

func main() {
	var departmentNum int

	_, err := fmt.Scan(&departmentNum)
	if err != nil {
		return
	}

	for range departmentNum {
		var workerNum int

		_, err := fmt.Scan(&workerNum)
		if err != nil {
			return
		}

		currentTemperature := NewTemperature(minTemperature, maxTemperature)

		for range workerNum {
			var (
				preferredTemperature int
				operand              string
			)

			_, err := fmt.Scan(&operand, &preferredTemperature)
			if err != nil {
				return
			}

			result := currentTemperature.getSuitableTemperature(operand, preferredTemperature)
			fmt.Println(result)
		}
	}
}
