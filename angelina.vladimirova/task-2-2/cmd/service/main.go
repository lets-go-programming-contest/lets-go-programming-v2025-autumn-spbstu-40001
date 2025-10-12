package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
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
		n int
		k int
	)

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	ratings := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&ratings[i])
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

	if k < 1 || k > n {
		fmt.Println("There is no such dish")
		return
	}

	h := &MinHeap{}
	heap.Init(h)

	for _, rating := range ratings {
		heap.Push(h, rating)
	}

	var result int
	for i := 0; i < k; i++ {
		result = heap.Pop(h).(int)
	}

	fmt.Println(result)
}
