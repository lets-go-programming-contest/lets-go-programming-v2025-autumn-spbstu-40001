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
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	for i := 0; i < n; i++ {
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())

		low := minTemp
		high := maxTemp

		for j := 0; j < k; j++ {
			sc.Scan()
			parts := strings.Split(sc.Text(), " ")
			t, _ := strconv.Atoi(parts[1])
			op := parts[0]
			fmt.Printf("%s %d\n", op, t)
		}
	}
}
