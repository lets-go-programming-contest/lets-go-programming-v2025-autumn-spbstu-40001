package main

import (
	"fmt"
)

func main() {
	var (
		n, k int
		var minTemp int
		var maxTemp int
		var optimalTemp int
		var border string
	)
	
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid departure count")
		return
	}
	
	for range n {
		_, err = fmt.Scan(&k)
		if err != nil {
			fmt.Println("Invalid empolyees count")
			return
		}
		
		minTemp = 15
		maxTemp = 30
		
		for range k {
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
			fmt.println(minTemp)
		} else {
			fmt.Println(-1)
		}
	}
}

