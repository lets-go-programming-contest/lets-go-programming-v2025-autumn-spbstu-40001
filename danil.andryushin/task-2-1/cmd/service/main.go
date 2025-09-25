package main

import "fmt"

func main() {
	var nDepartments uint

	_, err := fmt.Scan(&nDepartments)
	if err != nil || nDepartments == 0 {
		fmt.Println("Invalid departments count")

		return
	}

	for range nDepartments {
		var nEmployees uint

		_, err = fmt.Scan(&nEmployees)
		if err != nil || nEmployees == 0 {
			fmt.Println("Invalid employees count")

			return
		}

		var (
			minTemp, maxTemp uint = 15, 30
			currentTemp      uint
			operator         string
		)

		for range nEmployees {
			_, err := fmt.Scan(&operator, &currentTemp)
			if err != nil {
				fmt.Println("Failed to read employee wish")

				return
			}

			switch operator {
			case ">=":
				minTemp = max(minTemp, currentTemp)
			case "<=":
				maxTemp = min(maxTemp, currentTemp)
			default:
				fmt.Println("Unknown operator")

				return
			}

			if maxTemp >= minTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
