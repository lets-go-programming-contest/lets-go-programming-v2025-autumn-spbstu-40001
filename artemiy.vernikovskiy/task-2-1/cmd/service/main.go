package main

import (
	"fmt"
)

func main() {
	var (
		n, k int
		minTemp, maxTemp int
		optimalTemp int
		border string
	)
	
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Invalid departure count")
		return
	}
	
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&k)
		if err != nil || k <= 0 {
			fmt.Println("Invalid empolyees count")
			return
		}
		
		minTemp = 15
		maxTemp = 30
		optimalTemp = 0
		
		for j := 0; j < k; j++ {
			_, err = fmt.Scan(&border, &optimalTemp)
			if err != nil {
				fmt.Println("Invalid temperature")
				return
			}
			
			if border == ">=" {
				minTemp = max(minTemp, optimalTemp)
			} else if border == "<=" {
				maxTemp = min(maxTemp, optimalTemp)
			} else {
				fmt.Println("Wrong operator")
			}
		}
		
		if maxTemp >= minTemp {
			fmt.Println(minTemp)
		} else {
			fmt.Println(-1)
		}
	}
}

