package main

import (
	"errors"
	"fmt"
)

const (
	minComfortTemp = 15
	maxComfortTemp = 30
)

var ErrInvalidOperation = errors.New("invalid temperature operation")

type TemperatureRange struct {
	min int
	max int
}

func (tr *TemperatureRange) Update(operation string, temperature int) error {
	switch operation {
	case ">=":
		if temperature > tr.min {
			tr.min = temperature
		}
	case "<=":
		if temperature < tr.max {
			tr.max = temperature
		}
	default:
		return fmt.Errorf("operation '%s': %w", operation, ErrInvalidOperation)
	}

	return nil
}

func (tr *TemperatureRange) GetOptimal() int {
	if tr.min > tr.max {
		return -1
	}

	return tr.min
}

func main() {
	var departmentCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Failed to read department count:", err)

		return
	}

	for range departmentCount {
		var employeeCount int

		_, err := fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("Failed to read employee count:", err)

			return
		}

		tempRange := TemperatureRange{minComfortTemp, maxComfortTemp}

		for range employeeCount {
			var (
				operation   string
				temperature int
			)

			_, err := fmt.Scan(&operation, &temperature)
			if err != nil {
				fmt.Println("Failed to read temperature data:", err)

				return
			}

			err = tempRange.Update(operation, temperature)
			if err != nil {
				fmt.Println("Invalid temperature operation:", err)

				return
			}

			fmt.Println(tempRange.GetOptimal())
		}
	}
}
