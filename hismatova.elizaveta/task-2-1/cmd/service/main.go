package main

import (
	"fmt"
)

const (
	MinTemp = 15
	MaxTemp = 30
)

func main() {
	var departCount int
	if _, err := fmt.Scanln(&departCount); err != nil {
		fmt.Println("Error reading number of departments:", err)
		return
	}
	for i := 0; i < departCount; i++ {
		var peopleCount int
		if _, err := fmt.Scanln(&peopleCount); err != nil {
			fmt.Println("Error reading number of people:", err)
			return
		}
		minT, maxT := MinTemp, MaxTemp
		for j := 0; j < peopleCount; j++ {
			var operation string
			var needTemp int
			if _, err := fmt.Scanf("%s %d", &operation, &needTemp); err != nil {
				fmt.Println("Error reading operation and temperature:", err)
				return
			}
			switch operation {
			case "<=":
				if needTemp < maxT {
					maxT = needTemp
				}
			case ">=":
				if needTemp > minT {
					minT = needTemp
				}
			}
		}
		if minT > maxT {
			fmt.Println(-1)
		} else {
			fmt.Println(minT)
		}
	}
}
