package main

import (
	"errors"
	"fmt"
)

type TemperatureLimits struct {
	minTemperature int
	maxTemperature int
}

var ErrInvalidSymbol = errors.New("invalid symbol of set temperature limit")

func (limits *TemperatureLimits) processTemperature(border string, currentTemperature int) (int, error) {
	switch border {
	case ">=":
		limits.minTemperature = max(limits.minTemperature, currentTemperature)
	case "<=":
		limits.maxTemperature = min(limits.maxTemperature, currentTemperature)
	default:
		return 0, ErrInvalidSymbol
	}

	if limits.minTemperature <= limits.maxTemperature {
		return limits.minTemperature, nil
	}

	return -1, nil
}

func main() {
	var departmentsCount int
	if _, err := fmt.Scanln(&departmentsCount); err != nil {
		fmt.Println("failed to read departments count: ", err)

		return
	}

	for range departmentsCount {
		var (
			limits     = TemperatureLimits{15, 30}
			staffCount int
		)

		if _, err := fmt.Scanln(&staffCount); err != nil {
			fmt.Println("failed to read staff count: ", err)

			return
		}

		for range staffCount {
			var (
				border             string
				currentTemperature int
			)

			if _, err := fmt.Scanln(&border, &currentTemperature); err != nil {
				fmt.Println("failed to read current temperature limit: ", err)

				return
			}

			if value, err := limits.processTemperature(border, currentTemperature); err != nil {
				fmt.Println(err)

				return
			} else {
				fmt.Println(value)
			}
		}
	}
}
