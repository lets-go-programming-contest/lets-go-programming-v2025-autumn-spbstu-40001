package main

import "fmt"

type conditioner struct {
	defaultTemperature [2]int
	match              bool
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

		temperatureRange := conditioner{[2]int{15, 30}, true}

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
			switch direction {
			case ">=":
				if degrees <= temperatureRange.defaultTemperature[1] {
					temperatureRange.defaultTemperature[0] = degrees
				} else {
					temperatureRange.match = false
				}

			case "<=":
				if degrees >= temperatureRange.defaultTemperature[0] {
					temperatureRange.defaultTemperature[1] = degrees
				} else {
					temperatureRange.match = false
				}
			}
			if !temperatureRange.match {
				fmt.Println("-1")
			} else {
				fmt.Println(temperatureRange.defaultTemperature[0])
			}
		}
	}
}
