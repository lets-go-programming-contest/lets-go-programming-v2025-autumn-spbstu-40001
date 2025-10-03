package main

import "fmt"

func main() {
	var departmentsCount int
	if _, err := fmt.Scanln(&departmentsCount); err != nil {
		fmt.Println("Invalid input")

		return
	}

	for range departmentsCount {
		var (
			minTemperature = 15
			maxTemperature = 30
			staffCount     int
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

			func(border string, currentTemperature int) {
				if border == ">=" {
					minTemperature = max(minTemperature, currentTemperature)
				} else {
					maxTemperature = min(maxTemperature, currentTemperature)
				}

				if minTemperature <= maxTemperature {
					fmt.Println(minTemperature)
				} else {
					fmt.Println(-1)
				}
			}(border, currentTemperature)
		}
	}
}
