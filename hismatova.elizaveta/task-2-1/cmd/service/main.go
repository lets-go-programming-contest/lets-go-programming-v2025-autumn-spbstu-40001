package main

import "fmt"

const (
	MinTemp = 15
	MaxTemp = 30
)

func UpdateRangeAndGetTemp(
	minT, maxT int,
	operation string,
	temp int,
) (int, int) {
	switch operation {
	case "<=":
		if temp < maxT {
			maxT = temp
		}
	case ">=":
		if temp > minT {
			minT = temp
		}
	}
	if minT > maxT {
		return -1, maxT
	}

	return minT, maxT
}

func main() {
	var departCount int

	if _, err := fmt.Scanln(&departCount); err != nil {
		fmt.Println("Error reading number of departments:", err)
		return
	}

	for i := range departCount {
		var peopleCount int

		if _, err := fmt.Scanln(&peopleCount); err != nil {
			fmt.Println("Error reading number of people:", err)
			return
		}

		minT, maxT := MinTemp, MaxTemp

		for range peopleCount {
			var (
				operation string
				needTemp  int
			)

			if _, err := fmt.Scanf("%s %d\n", &operation, &needTemp); err != nil {
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

			if minT > maxT {
				fmt.Println(-1)
			} else {
				fmt.Println(minT)
			}
		}
	}
}
