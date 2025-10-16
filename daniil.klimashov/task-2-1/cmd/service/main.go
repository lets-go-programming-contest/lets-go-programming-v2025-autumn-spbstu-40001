package main

import "fmt"

const (
	defaultMinTemperature = 15
	defaultMaxTemperature = 30
)

type department struct {
	minTemp int
	maxTemp int
}

func newDepartment() *department {
	return &department{
		minTemp: defaultMinTemperature,
		maxTemp: defaultMaxTemperature,
	}
}

func (d *department) updateDesiredTemperature(op string, temp int) {
	switch op {
	case ">=":
		if temp > d.minTemp {
			d.minTemp = temp
		}
	case "<=":
		if temp < d.maxTemp {
			d.maxTemp = temp
		}
	}
}

func (d *department) getDesiredTemperature() int {
	if d.minTemp <= d.maxTemp {
		return d.minTemp
	}

	return -1
}

func main() {
	var (
		departmentAmount int
		employeeAmount   int
	)

	if _, err := fmt.Scan(&departmentAmount); err != nil {
		fmt.Println("Invalid number or departments")

		return
	}

	if _, err := fmt.Scan(&employeeAmount); err != nil {
		fmt.Println("Invalid number of employees")

		return
	}

	for range departmentAmount {
		dept := newDepartment()

		for range employeeAmount {
			var operator string
			if _, err := fmt.Scan(&operator); err != nil {
				fmt.Println("Invalid operator")

				return
			}

			var temp int
			if _, err := fmt.Scan(&temp); err != nil {
				fmt.Println("Invalid temperature limit value")

				return
			}

			dept.updateDesiredTemperature(operator, temp)

			desiredTemp := dept.getDesiredTemperature()
			fmt.Println(desiredTemp)
		}

		if _, err := fmt.Scan(&employeeAmount); err != nil {
			fmt.Println("Invalid number of employees")

			return
		}
	}
}
