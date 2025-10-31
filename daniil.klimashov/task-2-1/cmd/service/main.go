package main

import (
	"errors"
	"fmt"
)

const (
	defaultMinTemp = 15
	defaultMaxTemp = 30
)

var ErrInvalidOperator = errors.New("invalid operator")

type department struct {
	minTemp int
	maxTemp int
}

func newDepartment(minTemp, maxTemp int) *department {
	return &department{
		minTemp: minTemp,
		maxTemp: maxTemp,
	}
}

func (d *department) updateDesiredTemperature(operator string, temp int) error {
	switch operator {
	case ">=":
		if temp > d.minTemp {
			d.minTemp = temp
		}
	case "<=":
		if temp < d.maxTemp {
			d.maxTemp = temp
		}
	default:
		return fmt.Errorf("operation %q: %w", operator, ErrInvalidOperator)
	}

	return nil
}

func (d *department) getDesiredTemperature() int {
	if d.minTemp > d.maxTemp {
		return -1
	}

	return d.minTemp
}

func main() {
	var departmentAmount int

	if _, err := fmt.Scan(&departmentAmount); err != nil {
		fmt.Println("Invalid number or departments", err)

		return
	}

	for range departmentAmount {
		var employeeAmount int

		if _, err := fmt.Scan(&employeeAmount); err != nil {
			fmt.Println("Invalid number of employees", err)

			return
		}

		temperature := newDepartment(defaultMinTemp, defaultMaxTemp)

		for range employeeAmount {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator); err != nil {
				fmt.Println("Invalid operator", err)

				return
			}

			if _, err := fmt.Scan(&temp); err != nil {
				fmt.Println("Invalid temperature limit value", err)

				return
			}

			err := temperature.updateDesiredTemperature(operator, temp)
			if err != nil {
				fmt.Println("Error updating temperature:", err)

				return
			}

			fmt.Println(temperature.getDesiredTemperature())
		}
	}
}
