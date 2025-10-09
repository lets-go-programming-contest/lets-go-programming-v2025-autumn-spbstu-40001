package main

import "fmt"

const (
	minComfortTemp = 15
	maxComfortTemp = 30
)

func main() {
	var departmentCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Invalid input")

		return
	}

	for range departmentCount {
		var employeeCount int

		_, err := fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("Invalid input")

			return
		}

		minTemp := minComfortTemp
		maxTemp := maxComfortTemp

		for range employeeCount {
			var operation string
			var temperature int

			_, err := fmt.Scan(&operation, &temperature)
			if err != nil {
				fmt.Println("Invalid input")

				return
			}

			switch operation {
			case ">=":
				if temperature > minTemp {
					minTemp = temperature
				}
			case "<=":
				if temperature < maxTemp {
					maxTemp = temperature
				}
			}

			if minTemp > maxTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
