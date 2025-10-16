package main

import (
	"errors"
	"fmt"
)

type ComfortTemperature struct {
	minTemp int
	maxTemp int
}

var ErrInvalidSign error = errors.New("invalid sign")

func (temperature *ComfortTemperature) UpdateTemperature(sign string, value int) error {
	switch sign {
	case "<=":
		temperature.maxTemp = min(temperature.maxTemp, value)
	case ">=":
		temperature.minTemp = max(temperature.minTemp, value)
	default:
		return ErrInvalidSign
	}

	return nil
}

var ErrInvalidTemperature error = errors.New("invalid temperature")

func (temperature *ComfortTemperature) GetTemperature() (int, error) {
	if temperature.maxTemp < temperature.minTemp {
		return 0, ErrInvalidTemperature
	}

	return temperature.minTemp, nil
}

const minTemp, maxTemp, invalidTemp int = 15, 30, -1

func main() {
	var departmentsCount uint

	_, err := fmt.Scan(&departmentsCount)
	if err != nil {
		fmt.Println("invalid departments count:", err)

		return
	}

	for range departmentsCount {
		var employeesCount int

		_, err = fmt.Scan(&employeesCount)
		if err != nil {
			fmt.Println("invalid employee count:", err)

			return
		}

		temperature := ComfortTemperature{minTemp, maxTemp}

		for range employeesCount {
			var (
				conditionSign  string
				conditionValue int
			)

			_, err = fmt.Scan(&conditionSign, &conditionValue)
			if err != nil {
				fmt.Println("invalid temperature value:", err)

				return
			}

			err = temperature.UpdateTemperature(conditionSign, conditionValue)
			if err != nil {
				fmt.Println("failed to update temperature:", err)

				return
			}

			currentTemperature, err := temperature.GetTemperature()
			if err != nil {
				fmt.Println(invalidTemp)
			} else {
				fmt.Println(currentTemperature)
			}
		}
	}
}
