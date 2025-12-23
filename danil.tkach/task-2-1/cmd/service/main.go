package main

import (
	"errors"
	"fmt"
)

var ErrUndefinedOp = errors.New("undefined operation")

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

func (tr *TemperatureRange) Update(operation string, temp int) error {
	switch operation {
	case "<=":
		tr.maxT = min(tr.maxT, temp)
	case ">=":
		tr.minT = max(tr.minT, temp)
	default:
		return fmt.Errorf("%w: %s", ErrUndefinedOp, operation)
	}

	return nil
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
		fmt.Println("Failed to read count of departments:", err)

		return
	}

	for range departCount {
		var peopleCount int

		_, err := fmt.Scanln(&peopleCount)
		if err != nil {
			fmt.Println("Failed to read count of people:", err)

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
				fmt.Println("Failed to read operation or needed temperature:", err)

				return
			}

			updateErr := tempRange.Update(operation, needTemp)
			if updateErr != nil {
				fmt.Println("Failed while update temperature:", updateErr)

				return
			}

			fmt.Println(tempRange.GetOptimalTemp())
		}
	}
}
