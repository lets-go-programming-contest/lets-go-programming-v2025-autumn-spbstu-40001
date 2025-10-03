package main

import "fmt"

func main() {
	var departments_count int
	fmt.Scanln(&departments_count)
	for range departments_count {
		var min_temperature int = 15
		var max_temperature int = 30
		var staff_count int
		fmt.Scanln(&staff_count)
		for range staff_count {
			var border string
			var current_temperature int
			fmt.Scanln(&border, &current_temperature)
			func(border string, current_temperature int) {
				if border == ">=" {
					min_temperature = max(min_temperature, current_temperature)
				} else {
					max_temperature = min(max_temperature, current_temperature)
				}
				if min_temperature <= max_temperature {
					fmt.Println(min_temperature)
				} else {
					fmt.Println(-1)
				}
			}(border, current_temperature)
		}
	}
}
