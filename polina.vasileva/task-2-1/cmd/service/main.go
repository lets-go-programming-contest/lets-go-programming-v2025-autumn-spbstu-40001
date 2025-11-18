package main

import (
	"errors"
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30
)

var ErrInvalidOperation = errors.New("invalid operation")

type ConditionerT struct {
	minTemp int
	maxTemp int
}

func newConditioner(minTemperature, maxTemperature int) ConditionerT {
	return ConditionerT{
		minTemp: minTemperature,
		maxTemp: maxTemperature,
	}
}

func (cond *ConditionerT) changeTemp(sign string, degrees int) (int, error) {
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
		return 0, ErrInvalidOperation
	}

	if cond.minTemp <= cond.maxTemp {
		return cond.minTemp, nil
	}

	return -1, nil
}

func main() {
	var departNum int

	_, err := fmt.Scan(&departNum)
	if err != nil {
		fmt.Println("Failed to read count of departments", err)

		return
	}

	for range departNum {
		var emplCount int

		_, err := fmt.Scan(&emplCount)
		if err != nil {
			fmt.Println("Failed to read count of employees", err)

			return
		}

		conditioner := newConditioner(minTemp, maxTemp)

		for range emplCount {
			var sign string

			_, err = fmt.Scan(&sign)
			if err != nil {
				fmt.Println("Failed to read sign: ", err)

				return
			}

			var degrees int

			_, err = fmt.Scan(&degrees)
			if err != nil {
				fmt.Println("Failed to read temperature: ", err)

				return
			}

			result, err := conditioner.changeTemp(sign, degrees)
			if err != nil {
				fmt.Println("Failed to recalculate temp: ", err)

				return
			}

			fmt.Println(result)
		}
	}
}
