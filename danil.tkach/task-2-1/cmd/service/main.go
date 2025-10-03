package main

import (
	"fmt"
	"math"
)

const (
	MinTemp  = 15
	MaxTemp  = 30
	ErrorMin = 31
	ErrorMax = 14
)

func processEmployee(currentMin, currentMax int) (int, int) {
	var (
		operation string
		needTemp  int
	)

	_, err := fmt.Scanf("%s %d\n", &operation, &needTemp)
	if err != nil {
		return ErrorMin, ErrorMax
	}

	switch operation {
	case "<=":
		currentMax = int(math.Min(float64(currentMax), float64(needTemp)))
	case ">=":
		currentMin = int(math.Max(float64(currentMin), float64(needTemp)))
	default:
		return ErrorMin, ErrorMax
	}

	return currentMin, currentMax
}

func main() {
	var departCount int

	_, err := fmt.Scanln(&departCount)
	if err != nil {
		return
	}

	for range departCount {
		currentMin := MinTemp
		currentMax := MaxTemp

		var peopleCount int

		_, err := fmt.Scanln(&peopleCount)
		if err != nil {
			return
		}

		for range peopleCount {
			currentMin, currentMax = processEmployee(currentMin, currentMax)

			if currentMin <= currentMax && currentMax <= 30 && currentMin >= 15 {
				fmt.Println(currentMin)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
