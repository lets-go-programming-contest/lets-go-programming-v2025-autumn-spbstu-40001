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
	min      int
	max      int
	hasError bool
}

func NewOfficeThermostat(minTemp, maxTemp int) *OfficeThermostat {
	return &OfficeThermostat{
		min: minTemp,
		max: maxTemp,
	}
}

func (ot *OfficeThermostat) Process(operation string, temperature int) int {
	if ot.hasError {
		return -1
	}

	switch operation {
	case ">=":
		if temperature > ot.max {
			ot.hasError = true
			return -1
		}

		if temperature > ot.min {
			ot.min = temperature
		}
	case "<=":
		if temperature < ot.min {
			ot.hasError = true
			return -1
		}

		if temperature < ot.max {
			ot.max = temperature
		}
	}

	if ot.min > ot.max {
		ot.hasError = true
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

		thermostat := NewOfficeThermostat(minTemp, maxTemp)

		for range staffCount {
			var (
				operation   string
				temperature int
			)

			_, err := fmt.Scanf("%s %d\n", &operation, &temperature)
			if err != nil {
				log.Printf("Error reading operation and temperature: %v", err)
				os.Exit(1)
			}

			if operation != ">=" && operation != "<=" {
				fmt.Println(-1)

				continue
			}

			result := thermostat.Process(operation, temperature)
			fmt.Println(result)
		}
	}
}
