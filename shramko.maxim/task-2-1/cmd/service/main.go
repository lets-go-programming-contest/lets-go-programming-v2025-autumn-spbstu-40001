package main

import (
	"fmt"
	"math"
)

const (
	MinTemp = 15
	MaxTemp = 30
)

type TemperatureRange struct {
	minT int
	maxT int
}

func NewTemperatureRange(minTemp, maxTemp int) *TemperatureRange {
	return &TemperatureRange{
		minT: minTemp,
		maxT: maxTemp,
	}
}

func (tr *TemperatureRange) Update(operation string, temp int) {
	switch operation {
	case "<=":
		tr.maxT = int(math.Min(float64(tr.maxT), float64(temp)))
	case ">=":
		tr.minT = int(math.Max(float64(tr.minT), float64(temp)))
	default:
		tr.minT = tr.maxT + 1
	}
}

func (tr *TemperatureRange) GetOptimalTemp() int {
	if tr.minT > tr.maxT {
		return -1
	}

	return tr.minT
}

func main() {
	var departCount int

	_, err := fmt.Scanln(&departCount)
	if err != nil {
		return
	}

	for range departCount {
		var peopleCount int

		_, err := fmt.Scanln(&peopleCount)
		if err != nil {
			return
		}

		tempRange := NewTemperatureRange(MinTemp, MaxTemp)

		for range peopleCount {
			var (
				operation string
				needTemp  int
			)

			_, err := fmt.Scanf("%s %d\n", &operation, &needTemp)
			if err != nil {
				return
			}

			tempRange.Update(operation, needTemp)
			fmt.Println(tempRange.GetOptimalTemp())
		}
	}
}
