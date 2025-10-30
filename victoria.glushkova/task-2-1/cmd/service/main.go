package main

import (
	"fmt"
	"log"
	"os"
)

type OfficeThermostat struct {
	min int
	max int
}

func NewOfficeThermostat(min, max int) *OfficeThermostat {
	return &OfficeThermostat{
		min: min,
		max: max,
	}
}

func (ot *OfficeThermostat) Process(operation string, temperature int) int {
	switch operation {
	case ">=":
		if temperature > ot.min {
			ot.min = temperature
		}
	case "<=":
		if temperature < ot.max {
			ot.max = temperature
		}
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
		log.Printf("Error reading department count: %v", err)
		os.Exit(1)
	}

	for range departmentCount {
		var staffCount int

		_, err := fmt.Scan(&staffCount)
		if err != nil {
			log.Printf("Error reading staff count: %v", err)
			os.Exit(1)
		}

		thermostat := NewOfficeThermostat(15, 30)

		for range staffCount {
			var operation string
			var temperature int

			_, err := fmt.Scanf("%s %d\n", &operation, &temperature)
			if err != nil {
				log.Printf("Error reading operation and temperature: %v", err)
				os.Exit(1)
			}

			result := thermostat.Process(operation, temperature)
			fmt.Println(result)
		}
	}
}
