package main

import (
	"errors"
	"fmt"
)

var errInvalidDirection = errors.New("invalid direction")

type conditioner struct {
	minTemp int
	maxTemp int
}

func NewConditioner(minTemp int, maxTemp int) *conditioner {
	return &conditioner{
		minTemp: minTemp,
		maxTemp: maxTemp,
	}
}

func (c *conditioner) calculate(direction string, degrees int) error {
	match := true
	switch direction {
	case ">=":
		if degrees > c.maxTemp {
			match = false
		}

		if degrees >= c.minTemp {
			c.minTemp = degrees
		}

	case "<=":
		if degrees < c.minTemp {
			match = false
		}

		if degrees <= c.maxTemp {
			c.maxTemp = degrees
		}

	default:
		return errInvalidDirection
	}
	if !match {
		fmt.Println(-1)
	} else {
		fmt.Println(c.minTemp)
	}
	return nil
}

func main() {
	var departmentCount int

	if _, err := fmt.Scan(&departmentCount); err != nil {
		fmt.Println("invalid input:", err.Error())

		return
	}

	for range departmentCount {
		var employeeCount int

		if _, err := fmt.Scan(&employeeCount); err != nil {
			fmt.Println("invalid input:", err.Error())

			return
		}

		temperatureRange := NewConditioner(15, 30)

		for range employeeCount {
			var (
				direction string
				degrees   int
			)

			if _, err := fmt.Scan(&direction); err != nil {
				fmt.Println("invalid input:", err.Error())

				return
			}

			if _, err := fmt.Scan(&degrees); err != nil {
				fmt.Println("invalid input:", err.Error())

				return
			}

			if err := temperatureRange.calculate(direction, degrees); err != nil {
				fmt.Println("invalid input:", err.Error())
			}
		}
	}
}
