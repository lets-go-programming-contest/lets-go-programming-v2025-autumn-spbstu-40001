package main

import (
	"fmt"
)

func main() {
	var departNum uint

	_, err := fmt.Scan(&departNum)
	if err != nil {
		fmt.Println("Invalid input")

		return
	}

	for range departNum {
		var employNum int

		_, err := fmt.Scan(&employNum)
		if err != nil {
			fmt.Println("Invalid input")

			return
		}

		minTemperature := 15
		maxTemperature := 30

		for range employNum {
			var (
				condition          string
				currentTemperature int
			)

			_, err := fmt.Scan(&condition)
			if err != nil {
				fmt.Println("Invalid input")

				return
			}

			_, err = fmt.Scan(&currentTemperature)
			if err != nil {
				fmt.Println("Invalid input")

				return
			}

			switch condition {
			case "<=":
				maxTemperature = min(maxTemperature, currentTemperature)
			case ">=":
				minTemperature = max(minTemperature, currentTemperature)
			default:
				fmt.Println("Invalid condition")

				return
			}

			if minTemperature > maxTemperature {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemperature)
			}
		}
	}
}
