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
		recomT := 0

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
				if grade < minT {
					fmt.Println("-1")

					break innerFor
				}
				if maxT > grade {
					maxT = grade
				}
				recomT = minT
			case ">=":
				if grade > maxT {
					fmt.Println("-1")

					break innerFor
				}

				if minT < grade {
					minT = grade
				}
				recomT = maxT
			}

			fmt.Println(recomT)
		}
	}
}
