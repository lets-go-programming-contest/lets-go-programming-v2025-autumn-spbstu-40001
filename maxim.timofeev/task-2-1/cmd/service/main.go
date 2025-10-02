package main

import "fmt"

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

		temperatureRange := [2]int{15, 30}

	regulation:
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
				if degrees <= temperatureRange[1] {
					temperatureRange[0] = degrees
				}
			case "<=":
				if degrees >= temperatureRange[0] {
					temperatureRange[1] = degrees
				}
			default:
				fmt.Println("-1")

				break regulation
			}
			fmt.Println(temperatureRange[0])
		}
	}
}
