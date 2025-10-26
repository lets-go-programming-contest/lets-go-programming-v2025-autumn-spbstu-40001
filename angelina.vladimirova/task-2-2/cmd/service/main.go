package main

import (
	"container/heap"
	"errors"
	"fmt"
)

var (
	ErrInvalidCount    = errors.New("invalid count")
	ErrEmptyRatings    = errors.New("empty ratings")
	ErrHeapEmpty       = errors.New("heap is empty")
	ErrUnexpectedType  = errors.New("unexpected type from heap")
	ErrIndexOutOfRange = errors.New("index out of range")
)

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	if i < 0 || i >= h.Len() || j < 0 || j >= h.Len() {
		panic(ErrIndexOutOfRange)
	}

	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	if i < 0 || i >= h.Len() || j < 0 || j >= h.Len() {
		panic(ErrIndexOutOfRange)
	}

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
		panic(ErrHeapEmpty)
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func readCount() (int, error) {
	var count int

	_, err := fmt.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("read count: %w", err)
	}

	if count <= 0 {
		return 0, fmt.Errorf("%w: %d", ErrInvalidCount, count)
	}

	return count, nil
}

func readRatings(count int) ([]int, error) {
	ratings := make([]int, count)

	for index := range ratings {
		_, err := fmt.Scan(&ratings[index])
		if err != nil {
			return nil, fmt.Errorf("read rating: %w", err)
		}
	}

	return ratings, nil
}

func readPositionK() (int, error) {
	var positionK int

	_, err := fmt.Scan(&positionK)
	if err != nil {
		return 0, fmt.Errorf("read k: %w", err)
	}

	return positionK, nil
}

func findKthLargest(ratings []int, positionK int) (int, error) {
	if len(ratings) == 0 {
		return 0, ErrEmptyRatings
	}

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for _, rating := range ratings {
		heap.Push(maxHeap, rating)
	}

	if positionK > maxHeap.Len() {
		return 0, ErrHeapEmpty
	}

	for range positionK - 1 {
		heap.Pop(maxHeap)
	}

	result := heap.Pop(maxHeap)
	return result.(int), nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error: %v\n", r)
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

func readInput() (int, []int, int, error) {
	count, err := readCount()
	if err != nil {
		return 0, nil, 0, err
	}

	ratings, err := readRatings(count)
	if err != nil {
		return 0, nil, 0, err
	}

	positionK, err := readPositionK()
	if err != nil {
		return 0, nil, 0, err
	}

	return count, ratings, positionK, nil
}
