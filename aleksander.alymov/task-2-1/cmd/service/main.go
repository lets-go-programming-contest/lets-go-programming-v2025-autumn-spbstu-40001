package main

import (
	"errors"
	"fmt"
)

const (
	defaultLowerBound = 15
	defaultUpperBound = 30
	invalidTemp       = -1
)

var ErrInvalidOperation = errors.New("invalid operation")

type ClimateController struct {
	lowerBound int
	upperBound int
}

func NewClimateController() *ClimateController {
	return &ClimateController{
		lowerBound: defaultLowerBound,
		upperBound: defaultUpperBound,
	}
}

func (cc *ClimateController) ApplyConstraint(operator string, value int) error {
	switch operator {
	case ">=":
		if value > cc.lowerBound {
			cc.lowerBound = value
		}
	case "<=":
		if value < cc.upperBound {
			cc.upperBound = value
		}
	default:
		return fmt.Errorf("%w: %s", ErrInvalidOperation, operator)
	}

	return nil
}

func (cc *ClimateController) FindComfortableTemperature() int {
	if cc.lowerBound > cc.upperBound {
		return invalidTemp
	}

	return cc.lowerBound
}

func main() {
	var departCount int

	_, err := fmt.Scanln(&departCount)
	if err != nil {
		fmt.Println("Error reading number of departments:", err)

		return
	}

	for range departCount {
		var peopleCount int

		_, err := fmt.Scanln(&peopleCount)
		if err != nil {
			fmt.Println("Error reading number of people for department:", err)

			return
		}

		tempRange := NewClimateController()

		for range peopleCount {
			var (
				operation string
				needTemp  int
			)

			_, err := fmt.Scanf("%s %d\n", &operation, &needTemp)
			if err != nil {
				fmt.Println("Error reading operation or temperature:", err)

				return
			}

			constraintErr := tempRange.ApplyConstraint(operation, needTemp)
			if constraintErr != nil {
				fmt.Println("Error applying constraint:", constraintErr)

				return
			}

			fmt.Println(tempRange.FindComfortableTemperature())
		}
	}
}
