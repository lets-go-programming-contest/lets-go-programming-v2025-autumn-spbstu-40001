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
	if i < 0 || i >= len(*h) || j < 0 || j >= len(*h) {
		panic("index out of range")
	}

	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	if i < 0 || i >= len(*h) || j < 0 || j >= len(*h) {
		panic("index out of range")
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	v, ok := x.(int)
	if !ok {
		panic("IntHeap: Push received non-int value")
	}

	*h = append(*h, v)
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	length := len(old)

	if length == 0 {
		return nil
	}

	x := old[length-1]
	*h = old[:length-1]

	return x
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	var dishCount int

	_, err := fmt.Scan(&dishCount)
	if err != nil {
		fmt.Println("Fail in reading dish count:", err)

		return
	}

	ratings := &IntHeap{}
	heap.Init(ratings)

	for range dishCount {
		var rate int

		_, err := fmt.Scan(&rate)
		if err != nil {
			fmt.Println("Fail in reading rate:", err)

			return
		}

		heap.Push(ratings, rate)
	}

	var countK int

	_, err = fmt.Scan(&countK)
	if err != nil {
		fmt.Println("Fail in reading k:", err)

		return
	}

	if countK > ratings.Len() {
		fmt.Println("No such dish")

		return
	}

	for range countK - 1 {
		heap.Pop(ratings)
	}

	result := heap.Pop(ratings)
	if result == nil {
		fmt.Println("There is no such dish")

		return
	}

	fmt.Println(result)
}
