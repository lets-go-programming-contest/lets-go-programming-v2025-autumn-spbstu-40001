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
	min int
	max int
}

func NewOfficeThermostat() *OfficeThermostat {
	return &OfficeThermostat{
		min: minTemp,
		max: maxTemp,
	}
}

func (ot *OfficeThermostat) Process(operation string, temperature int) int {
	if operation == ">=" && temperature > ot.min {
		ot.min = temperature
	}

	if operation == "<=" && temperature < ot.max {
		ot.max = temperature
	}

	if ot.min > ot.max {
		return -1
	}

	return ot.min
}

func main() {
	var departmentCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}

	for range departmentCount {
		var staffCount int

		_, err := fmt.Scan(&staffCount)
		if err != nil {
			log.Printf("Error: %v", err)
			os.Exit(1)
		}

		thermostat := NewOfficeThermostat()

		for range staffCount {
			var operation string

			var temperature int

			_, err := fmt.Scanf("%s %d\n", &operation, &temperature)
			if err != nil {
				log.Printf("Error: %v", err)
				os.Exit(1)
			}

			result := thermostat.Process(operation, temperature)
			fmt.Println(result)
		}
	}
}
