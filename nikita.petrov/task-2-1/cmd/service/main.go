package main

import "fmt"

type TempManager struct {
	maxTemp int
	minTemp int
	optTemp int
}

func main() {
	var deptNum int

	fmt.Scan(&deptNum)

	for i := 0; i < deptNum; i++ {
		var (
			staffNum, wishfulTemp int
			condition             string
		)

		airConditioner := TempManager{30, 15, 0}

		fmt.Scan(&staffNum)

		for j := 0; j < staffNum; j++ {
			fmt.Scan(&condition)
			fmt.Scan(&wishfulTemp)

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
			}
		} else {
			someTM.minTemp = newTemp
			return false
		}
	default:
		return false
	}
	someTM.optTemp = someTM.minTemp
	return true
}
