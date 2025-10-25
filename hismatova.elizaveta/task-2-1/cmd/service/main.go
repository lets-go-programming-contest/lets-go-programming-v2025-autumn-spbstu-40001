package main

import "fmt"

const (
	MinTemp = 15
	MaxTemp = 30
)

type TemperatureRange struct {
	min int
	max int
}

func NewTemperatureRange() *TemperatureRange {
	return &TemperatureRange{
		min: MinTemp,
		max: MaxTemp,
	}
}

func (t *TemperatureRange) UpdateAndGet(operation string, temp int) int {
	switch operation {
	case "<=":
		if temp < t.max {
			t.max = temp
		}
	case ">=":
		if temp > t.min {
			t.min = temp
		}
	}

	if t.min > t.max {
		return -1
	}

	return t.min
}

func main() {
	var departCount int

	if _, err := fmt.Scanln(&departCount); err != nil {
		fmt.Println("Error reading number of departments:", err)

		return
	}

	for range departCount {
		var peopleCount int

		if _, err := fmt.Scanln(&peopleCount); err != nil {
			fmt.Println("Error reading number of people:", err)

			return
		}

		tempRange := NewTemperatureRange()

		for range peopleCount {
			var (
				operation string
				needTemp  int
			)

			if _, err := fmt.Scanf("%s %d\n", &operation, &needTemp); err != nil {
				fmt.Println("Error reading operation and temperature:", err)

				return
			}

			result := tempRange.UpdateAndGet(operation, needTemp)
			fmt.Println(result)
		}
	}
}
