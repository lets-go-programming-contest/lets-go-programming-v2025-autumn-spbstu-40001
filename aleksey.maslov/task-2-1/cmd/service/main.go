package main

import "fmt"

type ComfortZone struct {
	minTemp, maxTemp int
}

func (comf *ComfortZone) printComfortTemp(operation string, temp int) {
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
	lowerTempLimit = 15
	upperTempLimit = 30
)

func main() {
	var departmentCount, employeesCount int

	_, err := fmt.Scanln(&departmentCount)
	if err != nil {
		fmt.Println("Invalid input", err)

		return
	}

	for range departmentCount {
		_, err = fmt.Scanln(&employeesCount)
		if err != nil {
			fmt.Println("Invalid input", err)

			continue
		}

		comfortZone := ComfortZone{lowerTempLimit, upperTempLimit}

		for range employeesCount {
			var (
				operation string
				temp      int
			)

			_, err = fmt.Scanln(&operation, &temp)
			if err != nil {
				fmt.Println("Invalid input", err)

				continue
			}

			comfortZone.printComfortTemp(operation, temp)
		}
	}
}
