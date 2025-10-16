package main

import (
	"fmt"
	"log"
	"os"
)

const (
	minTemp = 15
	maxTemp = 30
)

type OfficeThermostat struct {
	minAllowed int
	maxAllowed int
}

func NewOfficeThermostat() *OfficeThermostat {
	return &OfficeThermostat{
		minAllowed: minTemp,
		maxAllowed: maxTemp,
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
	if _, err := fmt.Scan(&numDepartments); err != nil {
		log.Printf("Error reading number of departments: %v", err)
		os.Exit(1)
	}

	for range numDepartments {
		var staffSize int
		if _, err := fmt.Scan(&staffSize); err != nil {
			log.Printf("Error reading staff size: %v", err)
			os.Exit(1)
		}

		thermostat := NewOfficeThermostat()

		for range staffSize {
			var directive string
			var celsius int

			if _, err := fmt.Scanf("%s %d\n", &directive, &celsius); err != nil {
				log.Printf("Error reading temperature preference: %v", err)
				os.Exit(1)
			}

			optimalTemp := thermostat.ProcessPreference(directive, celsius)
			fmt.Println(optimalTemp)
		}
	}
}

