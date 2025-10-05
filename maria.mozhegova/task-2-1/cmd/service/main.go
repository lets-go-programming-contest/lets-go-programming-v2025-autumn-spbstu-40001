package main

import (
	"fmt"
)

func main() {
	var departNum uint
	_, err := fmt.Scan(&departNum)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	for range departNum {
		var employNum, resultTemperature int
		minTemperature, maxTemperature := 15, 30
		_, err := fmt.Scan(&employNum)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		for range employNum {
			var condition string
			_, err := fmt.Scan(&condition)
			if err != nil {
				fmt.Println("Invalid input")
				return
			}
			var currentTemperature int
			_, err = fmt.Scan(&currentTemperature)
			if err != nil {
				fmt.Println("Invalid input")
				return
			}
			switch condition {
			case "<=":
				maxTemperature = min(maxTemperature, currentTemperature)
				resultTemperature = min(maxTemperature, resultTemperature)
			case ">=":
				minTemperature = max(minTemperature, currentTemperature)
				resultTemperature = max(minTemperature, resultTemperature)
			default:
				fmt.Println("Invalid condition")
				return
			}
			if (resultTemperature < minTemperature) || (resultTemperature > maxTemperature) {
				resultTemperature = -1
			}
			fmt.Println(resultTemperature)
		}
	}
}
