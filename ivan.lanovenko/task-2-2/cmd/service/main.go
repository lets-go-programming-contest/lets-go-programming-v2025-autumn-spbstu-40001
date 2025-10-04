package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(value any) {
	num, ok := value.(int)
	if !ok {
		return
	}
	*h = append(*h, num)
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func printRating(h IntHeap, sequenceNumber int) {
	var value int
	for range sequenceNumber {

		num, ok := heap.Pop(&h).(int)
		if !ok {
			return
		}

		value = num
	}

	fmt.Println(value)
}

func main() {
	var numberOfDishes int

	if _, err := fmt.Scan(&numberOfDishes); err != nil {
		fmt.Println("[WRONG INPUT]")

		return
	}

	heapOfRatings := &IntHeap{}
	heap.Init(heapOfRatings)

	for range numberOfDishes {
		var current int

		if _, err := fmt.Scan(&current); err != nil {
			fmt.Println("[WRONG INPUT]")

			return
		}

		heap.Push(heapOfRatings, current)
	}

	var sequenceNumber int

	_, err := fmt.Scan(&sequenceNumber)

	if err != nil {
		fmt.Println("[WRONG INPUT]")

		return
	}

	if sequenceNumber > numberOfDishes {
		fmt.Println("[THE PRIORITY SEQUENCE NUMBER SHOULD NOT BE MORE THAN THE NUMBER OF DISHES]")

		return
	}

	printRating(*heapOfRatings, sequenceNumber)
}
