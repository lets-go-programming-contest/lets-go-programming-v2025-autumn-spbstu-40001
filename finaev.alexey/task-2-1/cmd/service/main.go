package main

import (
	"fmt"
)

type ComfortTemperature struct {
	minT int
	maxT int
}

func (temperature *ComfortTemperature) printTemperature(sign string, grade int) {
	switch sign {
	case "<=":
		temperature.maxT = min(temperature.maxT, grade)
	case ">=":
		temperature.minT = max(temperature.minT, grade)
	default:
		fmt.Println("Invalid input")

		return
	}

	if temperature.minT > temperature.maxT {
		fmt.Println(-1)
	} else {
		fmt.Println(temperature.minT)
	}
}

func main() {
	var depart int

	_, err := fmt.Scan(&depart)
	if err != nil {
		fmt.Println("Invalid:", err)

		return
	}

	for range depart {
		var (
			emploees    int
			temperature = ComfortTemperature{15, 30}
		)

		_, err = fmt.Scan(&emploees)
		if err != nil {
			fmt.Println("Invalid:", err)

			return
		}

		for range emploees {
			var (
				grade int
				sign  string
			)

			_, err = fmt.Scan(&sign)
			if err != nil {
				fmt.Println("Invalid:", err)

				return
			}

			_, err = fmt.Scan(&grade)
			if err != nil {
				fmt.Println("Invalid:", err)

				return
			}

			temperature.printTemperature(sign, grade)
		}
	}
}
