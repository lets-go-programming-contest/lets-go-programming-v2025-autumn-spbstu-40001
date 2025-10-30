package main

import (
	"errors"
	"fmt"
)

var errInvalidCondition = errors.New("invalid condition")

const (
	minTemp = 15
	maxTemp = 30
)

type Temperature struct {
	min int
	max int
}

func (temp *Temperature) changeTemp(cond string, currTemp int) error {
	switch cond {
	case "<=":
		temp.max = min(temp.max, currTemp)
	case ">=":
		temp.min = max(temp.min, currTemp)
	default:
		return errInvalidCondition
	}

	return nil
}

func (temp *Temperature) getOptimalTemp() int {
	if temp.min > temp.max {
		return -1
	}

	return temp.min
}

func main() {
	var departNum uint

	_, err := fmt.Scan(&departNum)
	if err != nil {
		fmt.Println("Failed to read count of departments:", err)

		return
	}

	for range departNum {
		var (
			employNum   int
			temperature = Temperature{minTemp, maxTemp}
		)

		_, err := fmt.Scan(&employNum)
		if err != nil {
			fmt.Println("Failed to read count of employees:", err)

			return
		}

		for range employNum {
			var (
				condition   string
				currentTemp int
			)

			_, err := fmt.Scan(&condition)
			if err != nil {
				fmt.Println("Failed to read condition:", err)

				return
			}

			_, err = fmt.Scan(&currentTemp)
			if err != nil {
				fmt.Println("Failed to read temperature:", err)

				return
			}

			err = temperature.changeTemp(condition, currentTemp)
			if err != nil {
				fmt.Println("Failed to change temperature:", err)

				return
			}

			fmt.Println(temperature.getOptimalTemp())
		}
	}
}
