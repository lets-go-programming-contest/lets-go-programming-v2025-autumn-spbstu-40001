package main

import (
	"fmt"
)

type ClimateControl struct {
	minTemp int
	maxTemp int
}

func NewClimateControl() *ClimateControl {
	return &ClimateControl{
		minTemp: 15,
		maxTemp: 30,
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

	for i := 0; i < departmentCount; i++ {
		var employeeCount int
		if _, err := fmt.Scanln(&employeeCount); err != nil {
			return
		}

		control := NewClimateControl()

		for j := 0; j < employeeCount; j++ {
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
