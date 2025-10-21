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

type MaxHeap struct {
	data []int
}

func (h *MaxHeap) Len() int {
	return len(h.data)
}

func (h *MaxHeap) Less(i, j int) bool {
	if i >= len(h.data) || j >= len(h.data) || i < 0 || j < 0 {
		return false
	}
	return h.data[i] > h.data[j]
}

func (h *MaxHeap) Swap(i, j int) {
	if i >= len(h.data) || j >= len(h.data) || i < 0 || j < 0 {
		return
	}
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MaxHeap) Push(x interface{}) {
	value, ok := x.(int)
	if !ok {
		return
	}
	h.data = append(h.data, value)
}

func (h *MaxHeap) Pop() interface{} {
	length := len(h.data)
	if length == 0 {
		return -1
	}
	x := h.data[length-1]
	h.data = h.data[0 : length-1]
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

func findKthLargest(ratings []int, positionK int) (int, error) {
	if len(ratings) == 0 {
		return -1, ErrEmptyRatings
	}

	maxHeap := &MaxHeap{data: ratings}
	heap.Init(maxHeap)

	var result int
	found := false

	for range positionK {
		value := heap.Pop(maxHeap)
		if value == -1 {
			return -1, ErrHeapEmpty
		}
		result = value.(int)
		found = true
	}

	if !found {
		return -1, ErrHeapEmpty
	}

	return result, nil
}

func main() {
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

	if result == -1 {
		fmt.Println("There is no such dish")
		return
	}

	fmt.Println(result)
}
