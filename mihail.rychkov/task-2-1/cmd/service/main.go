package main

import "fmt"

func main() {
	var nDepartments int

	_, err := fmt.Scan(&nDepartments)
	if err != nil {
		fmt.Println("Failed to read departments count")
		fmt.Println(err)

		return
	}

	for range nDepartments {
		var nEmployees int

		_, err = fmt.Scan(&nEmployees)
		if err != nil {
			fmt.Println("Failed to read employees count")
			fmt.Println(err)

			return
		}

		minTemperature, maxTemperature := 15, 30

		for range nEmployees {
			var (
				sign        string
				temperature int
			)

			_, err = fmt.Scan(&sign, &temperature)
			if err != nil {
				fmt.Println("Failed to read employee's wish")
				fmt.Println(err)

				return
			}

			switch sign {
			case ">=":
				minTemperature = max(minTemperature, temperature)
			case "<=":
				maxTemperature = min(maxTemperature, temperature)
			default:
				fmt.Println("Unknown comparison sign")

				return
			}

			if minTemperature <= maxTemperature {
				fmt.Println(minTemperature)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
