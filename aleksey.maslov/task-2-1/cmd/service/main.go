package main

import (
	"errors"
	"fmt"
)

var errInvalidOperation = errors.New("invalid operation")

type ComfortZone struct {
	minTemp, maxTemp int
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
		return errInvalidOperation
	}

	return nil
}

func (comf *ComfortZone) getComfortTemp() int {
	if comf.minTemp > comf.maxTemp {
		return -1
	} else {
		return comf.minTemp
	}
}

const (
	lowerTempLimit = 15
	upperTempLimit = 30
)

func main() {
	var departmentCount, employeesCount int

	_, err := fmt.Scanln(&departmentCount)
	if err != nil {
		fmt.Println("Failed to read department count", err)

		return
	}

	for range departmentCount {
		_, err = fmt.Scanln(&employeesCount)
		if err != nil {
			fmt.Println("Failed to read employees count", err)

			return
		}

		comfortZone := ComfortZone{lowerTempLimit, upperTempLimit}

		for range employeesCount {
			var operation string

			_, err = fmt.Scan(&operation)
			if err != nil {
				fmt.Println("Failed to read operation", err)

				return
			}

			var temp int

			_, err = fmt.Scan(&temp)
			if err != nil {
				fmt.Println("Failed to read temperature", err)

				return
			}

			err := comfortZone.updateTemp(operation, temp)
			if err != nil {
				fmt.Println("Failed to update temperature", err)

				return
			}

			println(comfortZone.getComfortTemp())
		}
	}
}
