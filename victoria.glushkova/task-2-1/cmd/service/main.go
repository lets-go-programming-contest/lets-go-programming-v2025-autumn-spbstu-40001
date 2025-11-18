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

func (ot *OfficeThermostat) Process(operation string, temperature int) int {
	switch operation {
	case ">=":
		if temperature > ot.max {
			return -1
		}
		if temperature > ot.min {
			ot.min = temperature
		}
	case "<=":
		if temperature < ot.min {
			return -1
		}
		if temperature < ot.max {
			ot.max = temperature
		}
	default:
		return -1
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

	for i := 0; i < departmentCount; i++ {
		var staffCount int

		_, err := fmt.Scan(&staffCount)
		if err != nil {
			log.Printf("Error reading staff count: %v", err)
			os.Exit(1)
		}

		thermostat := NewOfficeThermostat(minTemp, maxTemp)
		hasError := false

		for j := 0; j < staffCount; j++ {
			var (
				operation   string
				temperature int
			)

			_, err := fmt.Scanf("%s %d\n", &operation, &temperature)
			if err != nil {
				log.Printf("Error reading operation and temperature: %v", err)
				os.Exit(1)
			}

			if hasError {
				continue
			}

			result := thermostat.Process(operation, temperature)
			if result == -1 {
				hasError = true
				fmt.Println(-1)
			} else if j == staffCount-1 {
				fmt.Println(result)
			}
		}
	}
}
