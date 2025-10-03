package main

import "fmt"

type TemperatureLimits struct {
	minTemperature int
	maxTemperature int
}

func processTemperature(limits *TemperatureLimits, border string, currentTemperature int) {
	if border == ">=" {
		limits.minTemperature = max(limits.minTemperature, currentTemperature)
	} else {
		limits.maxTemperature = min(limits.maxTemperature, currentTemperature)
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
		fmt.Println("Invalid input")

		return
	}

	for range departmentsCount {
		var (
			limits     = TemperatureLimits{15, 30}
			staffCount int
		)

		if _, err := fmt.Scanln(&staffCount); err != nil {
			fmt.Println("Invalid input")

			return
		}

		for range staffCount {
			var (
				border             string
				currentTemperature int
			)

			if _, err := fmt.Scanln(&border, &currentTemperature); err != nil {
				fmt.Println("Invalid input")

				return
			}

			processTemperature(&limits, border, currentTemperature)
		}
	}
}
