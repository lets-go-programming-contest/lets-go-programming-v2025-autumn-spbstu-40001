package main

import (
	"errors"
	"fmt"
)

const (
	minT = 15
	maxT = 30
)

type ComfortTemperature struct {
	minT int
	maxT int
}

func NewTemperatureObject(minTemp, maxTemp int) ComfortTemperature {
	return ComfortTemperature{minT: minTemp, maxT: maxTemp}
}

var errSign = errors.New("unacceptable sign")

func (temperature *ComfortTemperature) CalculationTemperature(sign string, grade int) error {
	switch sign {
	case "<=":
		temperature.maxT = min(temperature.maxT, grade)
	case ">=":
		temperature.minT = max(temperature.minT, grade)
	default:
		return errSign
	}

	return nil
}

func (temperature *ComfortTemperature) ComfortTemperature() int {
	if temperature.minT > temperature.maxT {
		return -1
	}

	return temperature.minT
}

func main() {
	var depart int

	_, err := fmt.Scan(&depart)
	if err != nil {
		fmt.Println("Error reading the number of departments:", err)

		return
	}

	for range depart {
		var emploees int

		_, err = fmt.Scan(&emploees)
		if err != nil {
			fmt.Println("Error reading the number of emploees:", err)

			return
		}

		temperature := NewTemperatureObject(minT, maxT)

		for range emploees {
			var (
				grade int
				sign  string
			)

			_, err = fmt.Scan(&sign)
			if err != nil {
				fmt.Println("Error reading the sign:", err)

				return
			}

			_, err = fmt.Scan(&grade)
			if err != nil {
				fmt.Println("Error reading the grade:", err)

				return
			}

			err = temperature.CalculationTemperature(sign, grade)
			if err != nil {
				fmt.Println("Failed to calculate temperature:", err)

				return
			}

			fmt.Println(temperature.ComfortTemperature())
		}
	}
}
