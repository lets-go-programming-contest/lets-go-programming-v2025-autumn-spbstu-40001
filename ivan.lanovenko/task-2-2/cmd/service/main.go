package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(value any) {
	num, ok := value.(int)
	if !ok {
		panic(fmt.Sprintf("IntHeap.Push: expected int, got %T", value))
	}

	*h = append(*h, num)
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	if n == 0 {
		return nil
	}

	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	var numberOfDishes int

	if _, err := fmt.Scan(&numberOfDishes); err != nil {
		fmt.Println("Failed to read number of dishes: ", err)

		return
	}

	heapOfRatings := &IntHeap{}
	heap.Init(heapOfRatings)

	for range numberOfDishes {
		var current int

		if _, err := fmt.Scan(&current); err != nil {
			fmt.Println("Failed to read current rating: ", err)

			return
		}

		heap.Push(heapOfRatings, current)
	}

	var sequenceNumber int

	_, err := fmt.Scan(&sequenceNumber)
	if err != nil {
		fmt.Println("Failed to read sequenceNumber: ", err)

		return
	}

	if sequenceNumber > numberOfDishes {
		fmt.Println("The priority sequence number should not be more than the number of dishes")

		return
	}

	var value int

	for range sequenceNumber {
		num, ok := heap.Pop(heapOfRatings).(int)
		if !ok {
			return
		}

		value = num
	}

	fmt.Println(value)
}
