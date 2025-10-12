package main

import "fmt"

const (
	minTemperature = 15
	maxTemperature = 30
)

func main() {
	var numberOfTests int

	_, err := fmt.Scan(&numberOfTests)
	if err != nil {
		fmt.Println("Invalid input:", err)

		return
	}

	for range numberOfTests {
		var numberOfRequests int

		_, err := fmt.Scan(&numberOfRequests)
		if err != nil {
			fmt.Println("Invalid input:", err)

			return
		}

		currentMin := minTemperature
		currentMax := maxTemperature

		for range numberOfRequests {
			var (
				sign        string
				temperature int
			)

			_, err = fmt.Scan(&sign, &temperature)
			if err != nil {
				fmt.Println("Invalid input:", err)

				return
			}

			switch sign {
			case ">=":
				if temperature > currentMin {
					currentMin = temperature
				}
			case "<=":
				if temperature < currentMax {
					currentMax = temperature
				}
			}

			if currentMin > currentMax {
				fmt.Println(-1)
			} else {
				fmt.Println(currentMin)
			}
		}
	}
}
