package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	minTemp = 15
	maxTemp = 30
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for range make([]struct{}, n) {
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())

		low := minTemp
		high := maxTemp

		for range make([]struct{}, k) {
			scanner.Scan()
			parts := strings.Split(scanner.Text(), " ")
			temp, _ := strconv.Atoi(parts[1])

			if parts[0] == ">=" {
				if temp > low {
					low = temp
				}
			} else {
				if temp < high {
					high = temp
				}
			}

			if low > high {
				fmt.Println(-1)
			} else {
				fmt.Println(low)
			}
		}
	}
}
