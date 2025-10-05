package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x any) {
	if val, check := x.(int); check {
		*h = append(*h, val)
	} else {
		fmt.Println("invalid type in heap push")
	}
}

func (h *IntHeap) Pop() any {
	old := *h
	length := len(old)
	top := old[length-1]
	*h = old[:length-1]

	return top
}

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

	currentHeap := &IntHeap{}
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
