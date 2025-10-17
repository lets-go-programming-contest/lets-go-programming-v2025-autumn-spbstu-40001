package main

import (
	"fmt"

	"github.com/atroxxxxxx/task-2-1/internal/controller"
)

const (
	minTemp, maxTemp = 15, 30
	wrongTemp        = -1
)

func main() {
	var nDepartments uint

	_, err := fmt.Scan(&nDepartments)
	if err != nil {
		fmt.Println("invalid departments count:", err)

		return
	}

	for range nDepartments {
		var (
			nEmployees            uint
			temperatureController = controller.New(minTemp, maxTemp)
		)

		_, err := fmt.Scan(&nEmployees)
		if err != nil {
			fmt.Println("invalid employees count:", err)

			return
		}

		for range nEmployees {
			var desire controller.Desire

			_, err := fmt.Scan(&desire.Sign, &desire.DesiredTemperature)
			if err != nil {
				fmt.Println("invalid desire format:", err)

				return
			}

			err = temperatureController.ChangeTemperature(desire)
			if err != nil {
				fmt.Println("failed to change temperature:", err)

				return
			}

			temperature, err := temperatureController.GetTemperature()
			if err != nil {
				fmt.Println(wrongTemp)
			} else {
				fmt.Println(temperature)
			}
		}
	}
}
