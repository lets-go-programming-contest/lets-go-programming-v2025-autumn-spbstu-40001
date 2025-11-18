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

func NewOfficeThermostat(minTemp, maxTemp int) *OfficeThermostat {
	return &OfficeThermostat{
		min: minTemp,
		max: maxTemp,
	}
}

func (ot *OfficeThermostat) Process(operation string, temperature int) bool {
	switch operation {
	case ">=":
		if temperature > ot.max {
			return false
		}

		if temperature > ot.min {
			ot.min = temperature
		}
	case "<=":
		if temperature < ot.min {
			return false
		}

		if temperature < ot.max {
			ot.max = temperature
		}
	default:
		return false
	}

	return ot.min <= ot.max
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
		valid := true

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

			if valid {
				valid = thermostat.Process(operation, temperature)
			}
		}

		if valid {
			fmt.Println(thermostat.min)
		} else {
			fmt.Println(-1)
		}
	}
}
