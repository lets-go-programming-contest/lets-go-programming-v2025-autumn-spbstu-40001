package main

import (
	"errors"
	"fmt"
)

const (
	defaultMinTemp = 15
	defaultMaxTemp = 30
)

var errInvalidDirection = errors.New("invalid direction")

type conditioner struct {
	minTemp int
	maxTemp int
}

func newConditioner(minTemp int, maxTemp int) *conditioner {
	return &conditioner{
		minTemp: minTemp,
		maxTemp: maxTemp,
	}
}

func (c *conditioner) calculate(direction string, degrees int) error {
	switch direction {
	case ">=":
		if degrees >= c.minTemp {
			c.minTemp = degrees
		}

	case "<=":
		if degrees <= c.maxTemp {
			c.maxTemp = degrees
		}

	default:
		return errInvalidDirection
	}

	return nil
}

func (c *conditioner) getOptimalTemperature() int {
	if c.minTemp > c.maxTemp {
		return -1
	} else {
		return c.minTemp
	}
}

func main() {
	var departmentCount int

	if _, err := fmt.Scan(&departmentCount); err != nil {
		fmt.Println("failed to read count of departments:", err.Error())

		return
	}

	for range departmentCount {
		var employeeCount int

		if _, err := fmt.Scan(&employeeCount); err != nil {
			fmt.Println("failed to read count of employees:", err.Error())

			return
		}

		temperatureRange := newConditioner(defaultMinTemp, defaultMaxTemp)

		for range employeeCount {
			var (
				direction string
				degrees   int
			)

			if _, err := fmt.Scan(&direction); err != nil {
				fmt.Println("failed to read temperature direction:", err.Error())

				return
			}

			if _, err := fmt.Scan(&degrees); err != nil {
				fmt.Println("failed to read degrees:", err.Error())

				return
			}

			if err := temperatureRange.calculate(direction, degrees); err != nil {
				fmt.Println("failed to calculate optimal temperature:", err.Error())
			}

			fmt.Println(temperatureRange.getOptimalTemperature())
		}
	}
}
