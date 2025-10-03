package main

import "fmt"

func main() {
	var departmentCount, employeesCount int
	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("ERROR: input department count")

		return
	}
	for range departmentCount {
		_, err := fmt.Scanln(&employeesCount)
		if err != nil {
			fmt.Println("ERROR: input employees count")

			continue
		}
		processDepartment(employeesCount)
	}
}

func processDepartment(employeesCount int) {
	minTemperature := 15
	maxTemperature := 30
	var (
		operation   string
		temperature int
	)
	for range employeesCount {
		_, err := fmt.Scan(&operation)
		if err != nil {
			fmt.Println("ERROR: input operation")

			continue
		}
		_, err = fmt.Scan(&temperature)
		if err != nil {
			fmt.Println("ERROR: input temperature")

			continue
		}
		switch operation {
		case ">=":
			if temperature > minTemperature {
				minTemperature = temperature
			}
		case "<=":
			if temperature < maxTemperature {
				maxTemperature = temperature
			}
		default:
			fmt.Println("ERROR: invalid operation")

			continue
		}
		if minTemperature > maxTemperature {
			fmt.Println("-1")
		} else {
			fmt.Println(minTemperature)
		}
	}
}
