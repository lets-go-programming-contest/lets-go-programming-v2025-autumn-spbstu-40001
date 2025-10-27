package main

import (
	"container/heap"
	"fmt"

	"github.com/PigoDog/task-2-2/Internal/container/intheap"
)

func main() {
	var (
		dishCount int
		priority  int
	)

	if _, err := fmt.Scan(&dishCount); err != nil {
		fmt.Println("failed to read dishCount: ", err.Error())
	}

	currentHeap := &intheap.IntHeap{}
	heap.Init(currentHeap)

	for range dishCount {
		var currentPriority int

		if _, err := fmt.Scan(&currentPriority); err != nil {
			fmt.Println("failed to read currentPriority: ", err.Error())
		}

		heap.Push(currentHeap, currentPriority)
	}

	if _, err := fmt.Scan(&priority); err != nil {
		fmt.Println("failed to read priority: ", err.Error())
	}

	for range priority - 1 {
		heap.Pop(currentHeap)
	}

	todayDish := heap.Pop(currentHeap)

	if todayDish == nil {
		fmt.Println("no dish")

		return
	}

	fmt.Println(todayDish)
}
