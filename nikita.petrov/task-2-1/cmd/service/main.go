package main

import "fmt"

type TempManager struct {
	maxTemp int
	minTemp int
	optTemp int
	keepMin bool
}

func main() {
	var deptNum int

	fmt.Scan(&deptNum)

	for i := 0; i < deptNum; i++ {
		var (
			staffNum, wishfulTemp int
			condition             string
		)

		airConditioner := TempManager{30, 15, 0, false}

		fmt.Scan(&staffNum)
		fmt.Scan(&condition)
		fmt.Scan(&wishfulTemp)

		if !turnOnTempManager(&airConditioner, condition, wishfulTemp) {
			return
		}

		fmt.Println(airConditioner.optTemp)

		for j := 1; j < staffNum; j++ {
			fmt.Scan(&condition)
			fmt.Scan(&wishfulTemp)

			if !changeStatus(&airConditioner, condition, wishfulTemp) {
				return
			}

			fmt.Println(airConditioner.optTemp)
		}
	}
}

func turnOnTempManager(someTM *TempManager, condition string, initialTemp int) bool {
	if recalculateMinMax(someTM, condition, initialTemp) {
		if condition == ">=" {
			someTM.keepMin = true
		} else {
			someTM.keepMin = false
		}
		someTM.optTemp = initialTemp
		return true
	} else {
		return false
	}
}

func changeStatus(someTM *TempManager, condition string, newTemp int) bool {
	if recalculateMinMax(someTM, condition, newTemp) {
		if someTM.keepMin {
			someTM.optTemp = someTM.minTemp
		} else {
			someTM.optTemp = someTM.maxTemp
		}
		return true
	} else {
		return false
	}
}

func recalculateMinMax(someTM *TempManager, condition string, newTemp int) bool {
	switch condition {
	case "<=":
		if newTemp >= someTM.minTemp {
			if newTemp < someTM.maxTemp {
				someTM.maxTemp = newTemp
			}
		} else {
			fmt.Println(-1)
			return false
		}
	case ">=":
		if newTemp <= someTM.maxTemp {
			if newTemp > someTM.minTemp {
				someTM.minTemp = newTemp
			}
		} else {
			fmt.Println(-1)
			return false
		}
	default:
		fmt.Println(-1)
		return false
	}
	return true
}
