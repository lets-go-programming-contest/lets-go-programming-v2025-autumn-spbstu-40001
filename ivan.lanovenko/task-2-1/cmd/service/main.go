package main

import "fmt"

type TemperatureLimits struct {
	minTemperature int
	maxTemperature int
}

func (limits *TemperatureLimits) processTemperature(border string, currentTemperature int) {
	switch border {
	case ">=":
		limits.minTemperature = max(limits.minTemperature, currentTemperature)
	case "<=":
		limits.maxTemperature = min(limits.maxTemperature, currentTemperature)
	default:
		fmt.Println("Invalid symbol")

		return
	}

	if limits.minTemperature <= limits.maxTemperature {
		fmt.Println(limits.minTemperature)
	} else {
		fmt.Println(-1)
	}
}

func main() {
	var departmentsCount int
	if _, err := fmt.Scanln(&departmentsCount); err != nil {
		fmt.Println("Invalid input: ", err)

		return
	}

	for range departmentsCount {
		var (
			limits     = TemperatureLimits{15, 30}
			staffCount int
		)

		if _, err := fmt.Scanln(&staffCount); err != nil {
			fmt.Println("Invalid input: ", err)

			return
		}

		for range staffCount {
			var (
				border             string
				currentTemperature int
			)

			if _, err := fmt.Scanln(&border, &currentTemperature); err != nil {
				fmt.Println("Invalid input: ", err)

				return
			}

			limits.processTemperature(border, currentTemperature)
		}
	}
}
