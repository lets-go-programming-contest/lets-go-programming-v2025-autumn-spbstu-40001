package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	lenght := len(old)
	top := old[lenght-1]
	*h = old[:lenght-1]
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

	for i := 0; i < dishCount; i++ {
		var currentPriority int

		if _, err := fmt.Scan(&currentPriority); err != nil {
			fmt.Println("Invalid input")
		}

		arrayOfPriority[i] = currentPriority
	}

	if _, err := fmt.Scan(&priority); err != nil {
		fmt.Println("Invalid input")
	}

	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < dishCount; i++ {
		if h.Len() < priority {
			heap.Push(h, arrayOfPriority[i])
		} else if arrayOfPriority[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, arrayOfPriority[i])

		}
	}

	fmt.Println((*h)[0])

}
