package main

import "fmt"

type TempManager struct {
	maxTemp int
	minTemp int
	optTemp int
}

func main() {
	var deptNum int

	_, err := fmt.Scan(&deptNum)
	if err != nil {
		fmt.Println("Input error")

		return
	}

	for range deptNum {
		var (
			staffNum, wishfulTemp int
			condition             string
		)

		_, err := fmt.Scan(&staffNum)
		if err != nil {
			fmt.Println("Input error")

			return
		}

		airConditioner := TempManager{30, 15, 0}

		for range staffNum {
			_, err = fmt.Scan(&condition)
			if err != nil {
				fmt.Println("Input error")

				return
			}

			_, err = fmt.Scan(&wishfulTemp)
			if err != nil {
				fmt.Println("Input error")

				return
			}

			if changeStatus(&airConditioner, condition, wishfulTemp) {
				fmt.Println(airConditioner.optTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}

func changeStatus(someTM *TempManager, condition string, newTemp int) bool {
	switch condition {
	case "<=":
		if newTemp >= someTM.minTemp {
			if newTemp < someTM.maxTemp {
				someTM.maxTemp = newTemp
			}
		} else {
			someTM.maxTemp = newTemp

			return false
		}
	case ">=":
		if newTemp <= someTM.maxTemp {
			if newTemp > someTM.minTemp {
				someTM.minTemp = newTemp
				someTM.optTemp = someTM.minTemp
			}
		} else {
			someTM.minTemp = newTemp

			return false
		}
	}

	return true
}
