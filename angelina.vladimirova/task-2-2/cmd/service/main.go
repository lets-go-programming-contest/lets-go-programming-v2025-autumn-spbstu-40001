package main

import (
	"container/heap"
	"errors"
	"fmt"
)

var (
	ErrInvalidCount   = errors.New("invalid count")
	ErrEmptyRatings   = errors.New("empty ratings")
	ErrHeapEmpty      = errors.New("heap is empty")
	ErrUnexpectedType = errors.New("unexpected type from heap")
)

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	if i < 0 || i >= h.Len() || j < 0 || j >= h.Len() {
		panic("index out of range in maxheap")
	}

	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	if i < 0 || i >= h.Len() || j < 0 || j >= h.Len() {
		panic("index out of range in maxheap")
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	value, ok := x.(int)
	if !ok {
		panic("value is not an int")
	}

	*h = append(*h, value)
}

func (h *MaxHeap) Pop() interface{} {
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

	var count int

	_, err := fmt.Scan(&count)
	if err != nil {
		fmt.Println("Failed to read count:", err)

		return
	}

	ratings := &MaxHeap{}
	heap.Init(ratings)

	for range count {
		var rating int

		_, err := fmt.Scan(&rating)
		if err != nil {
			fmt.Println("Failed to read rating:", err)

			return
		}

		heap.Push(ratings, rating)
	}

	var positionK int

	_, err = fmt.Scan(&positionK)
	if err != nil {
		fmt.Println("Failed to read K:", err)

		return
	}

	if positionK > ratings.Len() {
		fmt.Println("There is no such dish")

		return
	}

	for range positionK - 1 {
		heap.Pop(ratings)
	}

	result := heap.Pop(ratings)
	if result == nil {
		fmt.Println("There is no such dish")

		return
	}

	fmt.Println(result)
}
