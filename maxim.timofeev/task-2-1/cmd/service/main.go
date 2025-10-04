package main

import (
	"errors"
	"fmt"
)

type conditioner struct {
	minTemp int
	maxTemp int
	match   bool
}

func (c *conditioner) directionManager(direction string, degrees int) error {
	switch direction {
	case ">=":
		if degrees > c.maxTemp {
			c.match = false
		}
		if degrees >= c.minTemp {
			c.minTemp = degrees
		}

	case "<=":
		if degrees < c.minTemp {
			c.match = false
		}
		if degrees <= c.maxTemp {
			c.maxTemp = degrees
		}

	default:
		return errors.New("invalid direction")
	}

	return nil
}

func main() {
	var departmentCount int

	if _, err := fmt.Scan(&departmentCount); err != nil {
		fmt.Println("invalid input")

		return
	}

	for range departmentCount {
		var employeeCount int

		if _, err := fmt.Scan(&employeeCount); err != nil {
			fmt.Println("invalid input")

			return
		}

		temperatureRange := conditioner{15, 30, true}

		for range employeeCount {
			var direction string

			var degrees int

			if _, err := fmt.Scan(&direction); err != nil {
				fmt.Println("invalid input")

				return
			}

			if _, err := fmt.Scan(&degrees); err != nil {
				fmt.Println("invalid input")

				return
			}

			if err := temperatureRange.directionManager(direction, degrees); err != nil {
				fmt.Println("invalid input")
			}

			if !temperatureRange.match {
				fmt.Println(-1)
			} else {
				fmt.Println(temperatureRange.minTemp)
			}
		}
	}
}
