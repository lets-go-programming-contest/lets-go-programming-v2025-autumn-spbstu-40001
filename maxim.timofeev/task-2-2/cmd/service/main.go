package main

import (
	"container/heap"
	"fmt"

	"github.com/PigoDog/task-2-2/package/container/intheap"
)

func main() {
	var (
		dishCount int
		priority  int
	)

	if _, err := fmt.Scan(&dishCount); err != nil {
		fmt.Println("Invalid input")
	}

	arrayOfPriority := make([]int, dishCount)

	for currentDish := range dishCount {
		var currentPriority int

		if _, err := fmt.Scan(&currentPriority); err != nil {
			fmt.Println("Invalid input")
		}

		arrayOfPriority[currentDish] = currentPriority
	}

	if _, err := fmt.Scan(&priority); err != nil {
		fmt.Println("Invalid input")
	}

	currentHeap := &intheap.IntHeap{}
	heap.Init(currentHeap)

	for currentDish := range dishCount {
		if currentHeap.Len() < priority {
			heap.Push(currentHeap, arrayOfPriority[currentDish])
		} else if arrayOfPriority[currentDish] > (*currentHeap)[0] {
			heap.Pop(currentHeap)
			heap.Push(currentHeap, arrayOfPriority[currentDish])
		}
	}

	fmt.Println((*currentHeap)[0])
}
