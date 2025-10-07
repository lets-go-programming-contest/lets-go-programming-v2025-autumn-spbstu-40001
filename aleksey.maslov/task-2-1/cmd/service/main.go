package main

import "fmt"

type ComfortZone struct {
	minTemp, maxTemp int
}

func (comf *ComfortZone) changeTemperature(operation string, temp int) {
	switch operation {
	case ">=":
		if temp > comf.minTemp {
			comf.minTemp = temp
		}
	case "<=":
		if temp < comf.maxTemp {
			comf.maxTemp = temp
		}
	default:
		fmt.Println("Invalid operation")
	}
	if comf.minTemp > comf.maxTemp {
		fmt.Println("-1")
	} else {
		fmt.Println(comf.minTemp)
	}
}

const (
	minTemperature = 15
	maxTemperature = 30
)

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
		comfortZone := ComfortZone{minTemperature, maxTemperature}
		for range employeesCount {
			var (
				operation string
				temp      int
			)

			_, err := fmt.Scan(&operation)
			if err != nil {
				fmt.Println("Invalid input", err)

				continue
			}
			_, err = fmt.Scan(&temp)
			if err != nil {
				fmt.Println("Invalid input", err)

				continue
			}
			comfortZone.changeTemperature(operation, temp)
		}
	}
}
