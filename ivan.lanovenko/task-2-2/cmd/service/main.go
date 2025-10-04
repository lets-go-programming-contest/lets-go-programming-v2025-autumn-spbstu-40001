package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (IntHeap IntHeap) Len() int           { return len(IntHeap) }
func (IntHeap IntHeap) Less(i, j int) bool { return IntHeap[i] > IntHeap[j] }
func (IntHeap IntHeap) Swap(i, j int)      { IntHeap[i], IntHeap[j] = IntHeap[j], IntHeap[i] }

func (IntHeap *IntHeap) Push(value any) {
	*IntHeap = append(*IntHeap, value.(int))
}

func (IntHeap *IntHeap) Pop() any {
	old := *IntHeap
	n := len(old)
	x := old[n-1]
	*IntHeap = old[0 : n-1]
	return x
}

func printRating(h IntHeap, sequenceNumber int) {
	var value int
	for range sequenceNumber {
		value = heap.Pop(&h).(int)
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
