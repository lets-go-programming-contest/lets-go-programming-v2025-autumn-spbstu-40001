package main

import (
	"fmt"
	"log"
	"os"
)

const minTemp = 15
const maxTemp = 30

type OfficeThermostat struct {
	min int
	max int
}

func NewOfficeThermostat() *OfficeThermostat {
	return &OfficeThermostat{min: minTemp, max: maxTemp}
}

func (ot *OfficeThermostat) Process(op string, temp int) int {
	if op == ">=" && temp > ot.min {
		ot.min = temp
	}
	if op == "<=" && temp < ot.max {
		ot.max = temp
	}
	if ot.min > ot.max {
		return -1
	}
	return ot.min
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
	for i := 0; i < n; i++ {
		var k int
		if _, err := fmt.Scan(&k); err != nil {
			log.Printf("Error: %v", err)
			os.Exit(1)
		}
		t := NewOfficeThermostat()
		for j := 0; j < k; j++ {
			var op string
			var temp int
			if _, err := fmt.Scanf("%s %d\n", &op, &temp); err != nil {
				log.Printf("Error: %v", err)
				os.Exit(1)
			}
			fmt.Println(t.Process(op, temp))
		}
	}
}
