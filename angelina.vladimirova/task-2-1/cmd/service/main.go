package main

import (
	"errors"
	"fmt"
)

const (
	lowerTempLimit = 15
	upperTempLimit = 30
)

var ErrUndefinedOperation = errors.New("undefined operation")

type ComfortZone struct {
	minTemp, maxTemp int
}

func NewComfortZone(minTemp, maxTemp int) *ComfortZone {
	return &ComfortZone{
		minTemp: minTemp,
		maxTemp: maxTemp,
	}
}

func (comf *ComfortZone) updateTemp(operation string, temp int) error {
	switch operation {
	case ">=":
		if temp > comf.minTemp {
			comf.minTemp = temp
		}
	case "<=":
		if temp < comf.maxTemp {
			comf.maxTemp = temp
		}
	default:
		return fmt.Errorf("%w: %s", ErrUndefinedOperation, operation)
	}

	return nil
}

func (comf *ComfortZone) getComfortTemp() int {
	if comf.minTemp > comf.maxTemp {
		return -1
	}

	return comf.minTemp
}

func main() {
	var departmentCount, employeesCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Failed to read department count:", err)

		return
	}

	for range departmentCount {
		_, err = fmt.Scan(&employeesCount)
		if err != nil {
			fmt.Println("Failed to read employees count:", err)

			return
		}

		comfortZone := NewComfortZone(lowerTempLimit, upperTempLimit)

		for range employeesCount {
			var (
				operation string
				temp      int
			)

			_, err = fmt.Scan(&operation, &temp)
			if err != nil {
				fmt.Println("Failed to read operation or temperature:", err)

				return
			}

			err = comfortZone.updateTemp(operation, temp)
			if err != nil {
				fmt.Println("Failed to update temperature:", err)

				return
			}

			fmt.Println(comfortZone.getComfortTemp())
		}
	}
}
