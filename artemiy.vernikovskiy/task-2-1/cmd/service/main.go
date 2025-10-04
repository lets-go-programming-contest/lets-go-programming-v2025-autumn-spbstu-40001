package main

import (
	"fmt"
)

func main() {
	var (
		nNumberOfOlimpic, kEmployeesCount int
		minTemp, maxTemp                  int
		optimalTemp                       int
		border                            string
	)

	_, err := fmt.Scan(&nNumberOfOlimpic)
	if err != nil || nNumberOfOlimpic < 0 {
		fmt.Println("Invalid departure count")

		return
	}

	for range nNumberOfOlimpic {
		_, err = fmt.Scan(&kEmployeesCount)
		if err != nil || kEmployeesCount < 0 {
			fmt.Println("Invalid employees count")

			return
		}

		minTemp = 15
		maxTemp = 30
		optimalTemp = 0

		for range kEmployeesCount {
			_, err = fmt.Scan(&border, &optimalTemp)
			if err != nil {
				fmt.Println("Invalid temperature")

				return
			}

			switch border {
			case ">=":
				minTemp = max(minTemp, optimalTemp)
			case "<=":
				maxTemp = min(maxTemp, optimalTemp)
			default:
				fmt.Println("Wrong operator")
			}

			if maxTemp >= minTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
