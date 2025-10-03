package main

import (
	"fmt"
	"math"
)

func processEmployee(currentMin, currentMax int) (int, int) {
	var (
		operation string
		need_temp int
	)
	_, err := fmt.Scanf("%s %d", &operation, &need_temp)
	if err != nil {
		return 31, 14
	}
	switch operation {
	case "<=":
		currentMax = int(math.Min(float64(currentMax), float64(need_temp)))
	case ">=":
		currentMax = int(math.Max(float64(currentMin), float64(need_temp)))
	default:
		return 31, 14
	}
	return currentMin, currentMax
}

func main() {
	var depart_count int
	_, err := fmt.Scanln(&depart_count)
	if err != nil {
		return
	}
	//currentMin := 15
	//currentMax := 30
	for count := 0; count < depart_count; count++ {
		var people_count int
		_, err := fmt.Scanln(&people_count)
		if err != nil {
			return
		}
		for count2 := 0; count2 < people_count; count2++ {
			processEmployee(0, 0)
		}
	}
	return
}
