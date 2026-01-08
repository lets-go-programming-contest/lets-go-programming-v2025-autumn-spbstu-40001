package main

import (
	"container/heap"
	"errors"
	"fmt"
)

var (
	ErrOfScan       = errors.New("scan failed")
	ErrInvalidRange = errors.New("value out of valid range")
)

// MaxHeap implements a max heap data structure.
type MaxHeap []int

// All methods use pointer receiver for consistency.
func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(firstIndex, secondIndex int) bool {
	heapSlice := *h
	if firstIndex < 0 || secondIndex < 0 || firstIndex >= len(heapSlice) || secondIndex >= len(heapSlice) {
		panic("index out of range in less")
	}

	return heapSlice[firstIndex] > heapSlice[secondIndex]
}

func (h *MaxHeap) Swap(firstIndex, secondIndex int) {
	heapSlice := *h
	if firstIndex < 0 || secondIndex < 0 || firstIndex >= len(heapSlice) || secondIndex >= len(heapSlice) {
		panic("index out of range in swap")
	}

	heapSlice[firstIndex], heapSlice[secondIndex] = heapSlice[secondIndex], heapSlice[firstIndex]
}

func (h *MaxHeap) Push(x any) {
	value, ok := x.(int)
	if !ok {
		panic("heap: Push of non-int value")
	}

	*h = append(*h, value)
}

func (h *MaxHeap) Pop() any {
	heapSlice := *h
	length := len(heapSlice)

	if length == 0 {
		return nil
	}

	last := heapSlice[length-1]
	*h = heapSlice[:length-1]

	return last
}

func readInt(prompt string, minValue, maxValue int) (int, error) {
	var value int
	_, err := fmt.Scan(&value)

	if err != nil {
		return 0, fmt.Errorf("invalid input of %s: %w", prompt, ErrOfScan)
	}

	if value < minValue || value > maxValue {
		return 0, fmt.Errorf("%s must be between %d and %d: %w", prompt, minValue, maxValue, ErrInvalidRange)
	}

	return value, nil
}

func main() {
	const (
		minDishes = 1
		maxDishes = 10000
		minRating = -10000
		maxRating = 10000
	)

	countOfDishes, err := readInt("count of dishes", minDishes, maxDishes)
	if err != nil {
		fmt.Println(err)

		return
	}

	dishHeap := &MaxHeap{}
	heap.Init(dishHeap)

	for range countOfDishes {
		rating, err := readInt("rating of dish", minRating, maxRating)
		if err != nil {
			fmt.Println(err)

			return
		}

		heap.Push(dishHeap, rating)
	}

	numOfPreference, err := readInt("num of preference", 1, countOfDishes)
	if err != nil {
		fmt.Println(err)

		return
	}

	if numOfPreference > 1 {
		for range numOfPreference - 1 {
			heap.Pop(dishHeap)
		}
	}

	kthLargestRaw := heap.Pop(dishHeap)
	kthLargest, ok := kthLargestRaw.(int)

	if !ok {
		fmt.Println("Type assertion failed")

		return
	}

	fmt.Println(kthLargest)
}
