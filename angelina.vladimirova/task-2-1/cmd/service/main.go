package main

import "fmt"

type ComfortZone struct {
	minTemp, maxTemp int
}

func (comf *ComfortZone) updateTemp(operation string, temp int) {
	switch operation {
	case ">=":
		if temp > comf.minTemp {
			comf.minTemp = temp
		}
	case "<=":
		if temp < comf.maxTemp {
			comf.maxTemp = temp
		}
	}
}

func (comf *ComfortZone) getComfortTemp() int {
	if comf.minTemp > comf.maxTemp {
		return -1
	}
	return comf.minTemp
}

const (
	lowerTempLimit = 15
	upperTempLimit = 30
)

func main() {
	var departmentCount, employeesCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	for range departmentCount {
		_, err = fmt.Scan(&employeesCount)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}

		comfortZone := ComfortZone{lowerTempLimit, upperTempLimit}

		for range employeesCount {
			var (
				operation string
				temp      int
			)

			_, err = fmt.Scan(&operation, &temp)
			if err != nil {
				fmt.Println("Invalid input")
				return
			}

			comfortZone.updateTemp(operation, temp)
			fmt.Println(comfortZone.getComfortTemp())
		}
	}
}
