package main

import "fmt"

func main() {
	var N int

	if _, err := fmt.Scan(&N); err != nil {
		fmt.Println("invalid input")
	}

	for range N {
		var K int

		if _, err := fmt.Scan(&K); err != nil {
			fmt.Println("invalid input")
		}

		temperatureRange := [2]int{15, 30}
		for range K {
			var direction string
			var degrees int

			if _, err := fmt.Scan(&direction); err != nil {
				fmt.Println("invalid input")
			}

			if _, err := fmt.Scan(&degrees); err != nil {
				fmt.Println("invalid input")
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

				break
			}
			fmt.Println(temperatureRange[0])
		}
	}
}
