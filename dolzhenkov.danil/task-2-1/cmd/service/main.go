package main

import (
	"fmt"
)

const (
	defaultMinTemp = 15
	defaultMaxTemp = 30
)

type ClimateControl struct {
	minTemp int
	maxTemp int
}

func NewClimateControl() *ClimateControl {
	return &ClimateControl{
		minTemp: defaultMinTemp,
		maxTemp: defaultMaxTemp,
	}
}

func (cc *ClimateControl) ApplyPreference(operator string, temperature int) {
	switch operator {
	case ">=":
		if temperature > cc.minTemp {
			cc.minTemp = temperature
		}
	case "<=":
		if temperature < cc.maxTemp {
			cc.maxTemp = temperature
		}
	}
}

func (cc *ClimateControl) CalculateComfort() int {
	if cc.minTemp > cc.maxTemp {
		return -1
	}

	return cc.minTemp
}

func main() {
	var departmentCount int
	if _, err := fmt.Scanln(&departmentCount); err != nil {
		return
	}

	for range departmentCount {
		var employeeCount int
		if _, err := fmt.Scanln(&employeeCount); err != nil {
			return
		}

		control := NewClimateControl()

		for range employeeCount {
			var (
				operator    string
				temperature int
			)

			if _, err := fmt.Scanf("%s %d\n", &operator, &temperature); err != nil {
				return
			}

			control.ApplyPreference(operator, temperature)
			fmt.Println(control.CalculateComfort())
		}
	}
}
