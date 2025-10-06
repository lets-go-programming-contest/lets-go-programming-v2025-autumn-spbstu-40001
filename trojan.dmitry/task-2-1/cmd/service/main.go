package main

import (
	"fmt"
)

func Min(first, second int) int {
	if first < second {
		return first
	}

	return second
}

func Max(first, second int) int {
	if first > second {
		return first
	}

	return second
}

func main() {
	var (
		department int
		workers    int
		num        int
		operator   string
	)

	_, err := fmt.Scan(&department)
	if err != nil {
		fmt.Println("Invalid input")

		return
	}

	for range department {
		_, err = fmt.Scan(&workers)
		if err != nil {
			fmt.Println("Invalid input")

			return
		}

		minLevel := 15
		maxLevel := 30

		for range workers {
			_, err = fmt.Scan(&operator)
			if err != nil {
				fmt.Println("Invalid input")

				return
			}

			_, err = fmt.Scan(&num)
			if err != nil {
				fmt.Println("Invalid input")

				return
			}

			switch operator {
			case ">=":
				minLevel = Max(minLevel, num)
			case "<=":
				maxLevel = Min(maxLevel, num)
			default:
				fmt.Printf("Invalid input")

				return
			}

			if minLevel <= maxLevel {
				fmt.Println(minLevel)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
