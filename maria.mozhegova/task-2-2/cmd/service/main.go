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

func (h *IntHeap) Push(x any) {
	value, ok := x.(int)
	if ok {
		*h = append(*h, value)
	} else {
		panic("value is not an int")
	}
}

func (h *IntHeap) Pop() any {
	if h.Len() == 0 {
		return nil
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	var dishNum uint

	_, err := fmt.Scan(&dishNum)
	if err != nil {
		fmt.Println("Invalid input:", err)

		return
	}

	ratings := &IntHeap{}
	heap.Init(ratings)

	for range dishNum {
		var rating int

		_, err := fmt.Scan(&rating)
		if err != nil {
			fmt.Println("Invalid input:", err)

			return
		}

		heap.Push(ratings, rating)
	}

	var dishK int

	_, err = fmt.Scan(&dishK)
	if err != nil {
		fmt.Println("Invalid input:", err)

		return
	}

	if dishK > ratings.Len() {
		fmt.Println("There is no such dish")

		return
	}

	for range dishK - 1 {
		heap.Pop(ratings)
	}

	fmt.Println(heap.Pop(ratings))
}
