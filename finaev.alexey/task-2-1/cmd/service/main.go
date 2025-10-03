package main

import (
	"fmt"
)

func main() {
	var (
		a    int
		sign string
		k    int
		n    int
	)
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	i := 1
	for i <= n {
		_, err = fmt.Scan(&k)
		if err != nil {
			fmt.Println("Invalid first operand")
			return
		}

		max_t := 30
		min_t := 15
		recom_t := 0
		j := 1

	innerFor:
		for j <= k {
			_, err = fmt.Scan(&sign)
			if err != nil {
				fmt.Println("Invalid first operand")
				return
			}
			_, err = fmt.Scan(&a)
			if err != nil {
				fmt.Println("Invalid first operand")
				return
			}

			switch sign {
			case "<=":
				if a < min_t {
					fmt.Println("-1")
					break innerFor
				}
				if max_t > a {
					max_t = a
				}
				recom_t = min_t
			case ">=":
				if a > max_t {
					fmt.Println("-1")
					break innerFor
				}
				if min_t < a {
					min_t = a
				}
				recom_t = max_t
			}

			fmt.Println(recom_t)
			j++
		}
		i++
	}
}
