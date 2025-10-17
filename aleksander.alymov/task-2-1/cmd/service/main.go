package main

import "fmt"

const (
	defaultLowerBound = 15
	defaultUpperBound = 30
	invalidTemp       = -1
)

type ClimateController struct {
	lowerBound int
	upperBound int
	valid      bool
}

func NewClimateController() *ClimateController {
	return &ClimateController{
		lowerBound: defaultLowerBound,
		upperBound: defaultUpperBound,
		valid:      true,
	}
}

func (cc *ClimateController) ApplyConstraint(operator string, value int) {
	if !cc.valid {
		return
	}

	switch operator {
	case ">=":
		if value > cc.lowerBound {
			cc.lowerBound = value
		}
	case "<=":
		if value < cc.upperBound {
			cc.upperBound = value
		}
	}

	if cc.lowerBound > cc.upperBound {
		cc.valid = false
	}
}

func (cc *ClimateController) FindComfortableTemperature() int {
	if !cc.valid {
		return invalidTemp
	}

	return cc.lowerBound
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

		tempRange := NewClimateController()

		for range peopleCount {
			var (
				operation string
				needTemp  int
			)

			_, err := fmt.Scanf("%s %d\n", &operation, &needTemp)
			if err != nil {
				return
			}

			tempRange.ApplyConstraint(operation, needTemp)
			fmt.Println(tempRange.FindComfortableTemperature())
		}
	}
}
