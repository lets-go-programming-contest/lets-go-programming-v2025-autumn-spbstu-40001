package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

type ConditionerT struct {
	minTemp int
	maxTemp int
}

func New(minTemperature, maxTemperature int) ConditionerT {
	return ConditionerT{
		minTemp: minTemperature,
		maxTemp: maxTemperature,
	}
}

func (cond *ConditionerT) changeTemp(sign string, degrees int) int {
	switch sign {
	case ">=":
		if degrees >= cond.minTemp {
			cond.minTemp = degrees
		}

	case "<=":
		if degrees <= cond.maxTemp {
			cond.maxTemp = degrees
		}
	default:
		fmt.Println("Invalid operation")
	}

	if cond.minTemp <= cond.maxTemp {
		return cond.minTemp
	} else {
		return -1
	}
}

func main() {
	var departNum int

	_, err := fmt.Scan(&departNum)
	if err != nil {
		fmt.Println("Invalid input", err)

		return
	}

	for range departNum {
		var emplCount int

		_, err := fmt.Scan(&emplCount)
		if err != nil {
			fmt.Println("Invalid input", err)

			return
		}

		conditioner := New(minTemp, maxTemp)

		for range emplCount {
			var sign string

			_, err = fmt.Scan(&sign)
			if err != nil {
				fmt.Println("Invalid input", err)

				return
			}

			var degrees int

			_, err = fmt.Scan(&degrees)
			if err != nil {
				fmt.Println("Invalid input", err)

				return
			}

			fmt.Println(conditioner.changeTemp(sign, degrees))
		}
	}
}
