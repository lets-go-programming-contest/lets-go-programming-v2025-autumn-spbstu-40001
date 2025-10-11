package main

import "fmt"

func main() {
	var (
		n int
	)

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	for i := 0; i < n; i++ {
		var (
			k    int
			sign string
			temp int
		)

		_, err := fmt.Scan(&k)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}

		minTemp := 15
		maxTemp := 30

		for j := 0; j < k; j++ {
			_, err = fmt.Scan(&sign, &temp)
			if err != nil {
				fmt.Println("Invalid input")
				return
			}

			switch sign {
			case ">=":
				if temp > minTemp {
					minTemp = temp
				}
			case "<=":
				if temp < maxTemp {
					maxTemp = temp
				}
			}

			if minTemp > maxTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
