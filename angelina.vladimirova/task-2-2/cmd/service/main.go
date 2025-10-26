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
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	value, ok := x.(int)
	if !ok {
		panic(ErrUnexpectedType)
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

func readInput() (int, []int, int, error) {
	var count int

	_, err := fmt.Scan(&count)
	if err != nil {
		return 0, nil, 0, fmt.Errorf("read count: %w", err)
	}

	if count <= 0 {
		return 0, nil, 0, fmt.Errorf("%w: %d", ErrInvalidCount, count)
	}

	ratings := make([]int, count)
	for index := range ratings {
		_, err := fmt.Scan(&ratings[index])
		if err != nil {
			return 0, nil, 0, fmt.Errorf("read rating: %w", err)
		}
	}

	var positionK int

	_, err = fmt.Scan(&positionK)
	if err != nil {
		return 0, nil, 0, fmt.Errorf("read k: %w", err)
	}

	return count, ratings, positionK, nil
}

func findKthLargest(ratings []int, positionK int) (int, error) {
	if len(ratings) == 0 {
		return 0, ErrEmptyRatings
	}

	maxHeap := MaxHeap(ratings)
	heap.Init(&maxHeap)

	if positionK > maxHeap.Len() {
		return 0, ErrHeapEmpty
	}

	for range positionK - 1 {
		heap.Pop(&maxHeap)
	}

	result := heap.Pop(&maxHeap)
	if result == nil {
		return 0, ErrHeapEmpty
	}

	return result.(int), nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	count, ratings, positionK, err := readInput()
	if err != nil {
		fmt.Printf("Invalid input: %v\n", err)

		return
	}

	if positionK < 1 || positionK > count {
		fmt.Println("There is no such dish")

		return
	}

	result, err := findKthLargest(ratings, positionK)
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		return
	}

	fmt.Println(result)
}
