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
	if i < 0 || i >= h.Len() || j < 0 || j >= h.Len() {
		panic("index out of range in intheap")
	}

	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	if i < 0 || i >= h.Len() || j < 0 || j >= h.Len() {
		panic("index out of range in inthea")
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	value, ok := x.(int)
	if !ok {
		panic("value is not an int")
	}

	*h = append(*h, value)
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic recovered:", r)
		}
	}()

	var dishNum uint

	_, err := fmt.Scan(&dishNum)
	if err != nil {
		fmt.Println("Failed to read count of dishes:", err)

		return
	}

	ratings := &IntHeap{}
	heap.Init(ratings)

	for range dishNum {
		var rating int

		_, err := fmt.Scan(&rating)
		if err != nil {
			fmt.Println("Failed to read dish rating:", err)

			return
		}

		heap.Push(ratings, rating)
	}

	var dishK int

	_, err = fmt.Scan(&dishK)
	if err != nil {
		fmt.Println("Failed to read K:", err)

		return
	}

	if dishK > ratings.Len() {
		fmt.Println("There is no such dish")

		return
	}

	for range dishK - 1 {
		heap.Pop(ratings)
	}

	result := heap.Pop(ratings)
	if result == nil {
		fmt.Println("There is no such dish")

		return
	}

	fmt.Println(result)
}
