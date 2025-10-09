package main

import "fmt"

const (
	minComfortTemp = 15
	maxComfortTemp = 30
)

type TemperatureRange struct {
	min int
	max int
}

func (tr *TemperatureRange) Update(operation string, temperature int) {
	switch operation {
	case ">=":
		if temperature > tr.min {
			tr.min = temperature
		}
	case "<=":
		if temperature < tr.max {
			tr.max = temperature
		}
	}
}

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

		tempRange := TemperatureRange{minComfortTemp, maxComfortTemp}

		for range employeeCount {
			var (
				operation   string
				temperature int
			)

			_, err := fmt.Scan(&operation, &temperature)
			if err != nil {
				fmt.Println("Invalid input")

				return
			}

			tempRange.Update(operation, temperature)

			if minComfortTemp > maxComfortTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minComfortTemp)
			}
		}
	}
}
