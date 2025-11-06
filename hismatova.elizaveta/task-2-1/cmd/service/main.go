package main

import (
	"errors"
	"fmt"
)

const (
	MinTemp = 15
	MaxTemp = 30
)

var ErrInvalidOperation = errors.New("invalid operation")

type TemperatureRange struct {
	min int
	max int
}

func NewTemperatureRange(minT, maxT int) *TemperatureRange {
	return &TemperatureRange{
		min: minT,
		max: maxT,
	}
}

func (t *TemperatureRange) UpdateAndGet(operation string, temp int) (int, error) {
	switch operation {
	case "<=":
		if temp < t.max {
			t.max = temp
		}
	case ">=":
		if temp > t.min {
			t.min = temp
		}
	default:
		return 0, fmt.Errorf("operation '%s': %w", operation, ErrInvalidOperation)
	}

	if t.min > t.max {
		return -1, nil
	}

	return t.min, nil
}

func main() {
	var departCount int

	if _, err := fmt.Scanln(&departCount); err != nil {
		fmt.Println("Error reading number of departments:", err)

		return
	}

	for range departCount {
		var peopleCount int

		if _, err := fmt.Scanln(&peopleCount); err != nil {
			fmt.Println("Error reading number of people:", err)

			return
		}

		tempRange := NewTemperatureRange(MinTemp, MaxTemp)

		for range peopleCount {
			var (
				operation string
				needTemp  int
			)

			if _, err := fmt.Scanf("%s %d\n", &operation, &needTemp); err != nil {
				fmt.Println("Error reading operation and temperature:", err)

				return
			}

			result, err := tempRange.UpdateAndGet(operation, needTemp)
			if err != nil {
				fmt.Println("Error:", err)

				return
			}

			fmt.Println(result)
		}
	}
}
