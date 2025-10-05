package main

import (
	"fmt"
)

func main() {
	var (
		grade    int
		sign     string
		emploees int
		depart   int
	)

	_, err := fmt.Scan(&depart)
	if err != nil {
		fmt.Println("Invalid input")

		return
	}

	for range depart {
		_, err = fmt.Scan(&emploees)
		if err != nil {
			fmt.Println("Invalid input")

			return
		}

		maxT := 30
		minT := 15

	innerFor:
		for range emploees {
			_, err = fmt.Scan(&sign)
			if err != nil {
				fmt.Println("Invalid input")

				return
			}

			_, err = fmt.Scan(&grade)
			if err != nil {
				fmt.Println("Invalid input")

				return
			}

			switch sign {
			case "<=":
				maxT = min(maxT, grade)
			case ">=":
				minT = max(minT, grade)
			default:
				fmt.Println("Invalid input")

				return
			}

			if minT > maxT {
				fmt.Println(-1)

				break innerFor
			} else {
				fmt.Println(minT)
			}
		}
	}
}
