package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var dishesNumber int

	_, err := fmt.Scan(&dishesNumber)
	if err != nil {
		fmt.Println("Parse error")

		return
	}

	ratingList := &IntHeap{}

	for range dishesNumber {
		var dishRating int
		_, err = fmt.Scan(&dishRating)
		if err != nil {
			fmt.Println("Parse error")

			return
		}

		heap.Push(ratingList, dishRating)
	}

	var wishedDish int

	_, err = fmt.Scan(&wishedDish)
	if err != nil {
		fmt.Println("Parse error")
	}

	for range wishedDish - 1 {
		heap.Pop(ratingList)
	}

	fmt.Println(heap.Pop(ratingList))
}
