package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	rating, ok := x.(int)
	if !ok {
		return
	}

	*h = append(*h, rating)
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	var totalDishes int

	_, err := fmt.Scan(&totalDishes)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	heapInstance := &IntHeap{}
	heap.Init(heapInstance)

	for range totalDishes {
		var rating int

		_, err := fmt.Scan(&rating)
		if err != nil {
			fmt.Println("Error:", err)

			return
		}

		heap.Push(heapInstance, rating)
	}

	var preferenceOrder int

	_, err = fmt.Scan(&preferenceOrder)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	temp := make(IntHeap, len(*heapInstance))
	copy(temp, *heapInstance)
	heap.Init(&temp)

	var result int

	for range preferenceOrder {
		item := heap.Pop(&temp)
		rating, ok := item.(int)

		if !ok {
			fmt.Println("Error: unexpected type")

			return
		}

		result = rating
	}

	fmt.Println(result)
}
