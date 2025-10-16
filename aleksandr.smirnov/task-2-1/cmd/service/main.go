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
	default:
		fmt.Println("Invalid operation")
	}
}

func (tr *TemperatureRange) GetOptimal() int {
	if tr.min > tr.max {
		return -1
	}

	return tr.min
}

func main() {
	var departmentCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Invalid input", err)

		return
	}

	for range departmentCount {
		var employeeCount int

		_, err := fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("Invalid input", err)

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
				fmt.Println("Invalid input", err)

				return
			}

			tempRange.Update(operation, temperature)
			fmt.Println(tempRange.GetOptimal())
		}
	}
}
