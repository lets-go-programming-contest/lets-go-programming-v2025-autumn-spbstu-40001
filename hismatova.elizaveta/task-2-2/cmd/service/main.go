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
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Error reading number of dishes:", err)

		return
	}

	arr := make([]int, n)
	for i := range arr {
		if _, err := fmt.Scan(&arr[i]); err != nil {
			fmt.Println("Error reading dish rating:", err)

			return
		}
	}

	var k int
	if _, err := fmt.Scan(&k); err != nil {
		fmt.Println("Error reading k:", err)

		return
	}

	h := &IntHeap{}
	heap.Init(h)

	for _, num := range arr {
		if h.Len() < k {
			heap.Push(h, num)
		} else {
			if num > (*h)[0] {
				heap.Pop(h)
				heap.Push(h, num)
			}
		}
	}

	fmt.Println((*h)[0])
}
