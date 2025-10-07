package main

import "fmt"

type 

func main() {
	var departmentCount, employeesCount int
	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Invalid input", err)

		return
	}
	for range departmentCount {
		_, err := fmt.Scanln(&employeesCount)
		if err != nil {
			fmt.Println("Invalid input", err)

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
			fmt.Println("Invalid input", err)

			continue
		}
		_, err = fmt.Scan(&temperature)
		if err != nil {
			fmt.Println("Invalid input", err)

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
			fmt.Println("Invalid operation", err)

			continue
		}
		if minTemperature > maxTemperature {
			fmt.Println("-1")
		} else {
			fmt.Println(minTemperature)
		}
	}
}
