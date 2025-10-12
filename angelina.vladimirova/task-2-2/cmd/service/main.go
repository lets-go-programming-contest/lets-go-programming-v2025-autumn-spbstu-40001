package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	value, ok := x.(int)
	if !ok {
		panic("value is not an int")
	}

	*h = append(*h, value)
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	var (
		count int
		k     int
	)

	_, err := fmt.Scan(&count)
	if err != nil {
		fmt.Println("Invalid input")

		return
	}

	ratings := make([]int, count)
	for index := range ratings {
		_, err := fmt.Scan(&ratings[index])
		if err != nil {
			fmt.Println("Invalid input")

			return
		}
	}

	_, err = fmt.Scan(&k)
	if err != nil {
		fmt.Println("Invalid input")

		return
	}

	if k < 1 || k > count {
		fmt.Println("There is no such dish")

		return
	}

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, rating := range ratings {
		heap.Push(minHeap, rating)
	}

	var result int
	for range k {
		value := heap.Pop(minHeap)
		result = value.(int)
	}

	fmt.Println(result)
}
