package main

import "fmt"

func main() {
	var departmentsCount int
	fmt.Scanln(&departmentsCount)
	for i := 0; i < departmentsCount; i++ {
		var minTemperature int = 15
		var maxTemperature int = 30
		var staffCount int
		fmt.Scanln(&staffCount)
		for i := 0; i < staffCount; i++ {
			var border string
			var currentTemperature int
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
