package main

import "fmt"

type OfficeThermostat struct {
	minAllowed int
	maxAllowed int
}

func NewOfficeThermostat() *OfficeThermostat {
	return &OfficeThermostat{
		minAllowed: 15,
		maxAllowed: 30,
	}
}

func (ot *OfficeThermostat) ProcessPreference(constraint string, degrees int) int {
	switch constraint {
	case ">=":
		if degrees > ot.minAllowed {
			ot.minAllowed = degrees
		}
	case "<=":
		if degrees < ot.maxAllowed {
			ot.maxAllowed = degrees
		}
	}

	if ot.minAllowed > ot.maxAllowed {
		return -1
	}
	return ot.minAllowed
}

func main() {
	var numDepartments int
	fmt.Scan(&numDepartments)

	for d := 0; d < numDepartments; d++ {
		var staffSize int
		fmt.Scan(&staffSize)

		thermostat := NewOfficeThermostat()

		for s := 0; s < staffSize; s++ {
			var directive string
			var celsius int
			fmt.Scanf("%s %d\n", &directive, &celsius)

			optimalTemp := thermostat.ProcessPreference(directive, celsius)
			fmt.Println(optimalTemp)
		}
	}
}
