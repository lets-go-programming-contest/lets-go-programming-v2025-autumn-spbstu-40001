package main

import (
	"container/heap"
	"errors"
	"fmt"
)

var (
	ErrHeapWrongType        = errors.New("attempt to add element of wrong type to heap")
	ErrHeapEmpty            = errors.New("attempt to extract element from empty heap")
	ErrPreferenceOutOfRange = errors.New("preference order out of range")
)

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
	value, ok := x.(int)
	if !ok {
		panic(ErrHeapWrongType)
	}

	*h = append(*h, value)
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	length := len(old)

	if length == 0 {
		panic(ErrHeapEmpty)
	}

	x := old[length-1]
	*h = old[0 : length-1]

	return x
}

func FindKthPreference(ratings []int, preferenceOrder int) (int, error) {
	if preferenceOrder < 1 || preferenceOrder > len(ratings) {
		return 0, fmt.Errorf("%w: %d is out of range [1, %d]",
			ErrPreferenceOutOfRange, preferenceOrder, len(ratings))
	}

	dishHeap := &MinHeap{}
	heap.Init(dishHeap)

	for _, rating := range ratings {
		heap.Push(dishHeap, rating)

		if dishHeap.Len() > preferenceOrder {
			heap.Pop(dishHeap)
		}
	}

	dish := heap.Pop(dishHeap)
	kthPreferenceDish, ok := dish.(int)
	if !ok {
		return 0, ErrHeapWrongType
	}

	return kthPreferenceDish, nil
}

func main() {
	var dishesCount int

	_, err := fmt.Scan(&dishesCount)
	if err != nil {
		fmt.Printf("Error reading dishes count: %v\n", err)
		return
	}

	ratings := make([]int, dishesCount)

	for dishIndex := range dishesCount {
		_, err = fmt.Scan(&ratings[dishIndex])
		if err != nil {
			fmt.Printf("Error reading dish rating: %v\n", err)
			return
		}
	}

	var preferenceOrder int

	_, err = fmt.Scan(&preferenceOrder)
	if err != nil {
		fmt.Printf("Error reading preference order: %v\n", err)
		return
	}

	kthPreference, err := FindKthPreference(ratings, preferenceOrder)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(kthPreference)
}
