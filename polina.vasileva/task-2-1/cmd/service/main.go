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

func (cond *ConditionerT) changeTemp(sign string, degrees int) {
	switch sign {
	case ">=":
		if degrees >= cond.minTemp {
			cond.minTemp = degrees
		}

	case "<=":
		if degrees <= cond.maxTemp {
			cond.maxTemp = degrees
		}
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

		conditioner := ConditionerT{minTemp, maxTemp}

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

			conditioner.changeTemp(sign, degrees)

			if conditioner.minTemp <= conditioner.maxTemp {
				fmt.Println(conditioner.minTemp)
			} else {
				fmt.Println("-1")
			}
		}
	}
}
