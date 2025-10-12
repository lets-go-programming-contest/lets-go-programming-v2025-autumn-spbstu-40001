package main

import "fmt"

func main() {
	var numberOfTests int

	_, err := fmt.Scan(&numberOfTests)
	if err != nil {
		fmt.Println("Invalid input")

		return
	}

	for testIndex := range numberOfTests {
		var (
			numberOfRequests int
			sign             string
			temperature      int
		)

		_, err := fmt.Scan(&numberOfRequests)
		if err != nil {
			fmt.Println("Invalid input")

			return
		}

		minTemperature := 15
		maxTemperature := 30

		for requestIndex := range numberOfRequests {
			_, err = fmt.Scan(&sign, &temperature)
			if err != nil {
				fmt.Println("Invalid input")

				return
			}

			switch sign {
			case ">=":
				if temperature > minTemperature {
					minTemperature = temperature
				}
			case "<=":
				if temperature < maxTemperature {
					maxTemperature = temperature
				}
			}

			if minTemperature > maxTemperature {
				fmt.Println(-1)

			} else {
				fmt.Println(minTemperature)
			}
		}
	}
}