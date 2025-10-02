package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N)
	fmt.Scan(&K)

	for i := 0; i < N; i++ {
		temperatureRange := [2]int{15, 30}
		for j := 0; j < K; j++ {
			var direction string
			var degrees int
			fmt.Scan(&direction)
			fmt.Scan(&degrees)
			if direction == ">=" && (degrees >= temperatureRange[0] && degrees <= temperatureRange[1]) {
				temperatureRange[0] = degrees
			} else if direction == "<=" && (degrees >= temperatureRange[0] && degrees <= temperatureRange[1]) {
				temperatureRange[1] = degrees
			} else {
				fmt.Println("-1")
			}
		}
	}
}
