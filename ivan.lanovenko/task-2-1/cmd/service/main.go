package main

import "fmt"

func main() {
	var departmentsCount int
	fmt.Scanln(&departmentsCount)

	for range departmentsCount {
		var (
			minTemperature = 15
			maxTemperature = 30
			staffCount     int
		)

		fmt.Scanln(&staffCount)

		for range staffCount {
			var (
				border             string
				currentTemperature int
			)

			fmt.Scanln(&border, &currentTemperature)
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
