package main

import (
	"container/heap"
	"errors"
	"fmt"
)

var ErrOfScan = errors.New("scan failed")

type MaxHeap []int

// Все методы используют pointer receiver для согласованности
func (iHeap *MaxHeap) Len() int {
	return len(*iHeap)
}

func (iHeap *MaxHeap) Less(firstIndex, secondIndex int) bool {
	h := *iHeap
	if firstIndex < 0 || secondIndex < 0 || firstIndex >= len(h) || secondIndex >= len(h) {
		panic("index out of range in less")
	}

	return h[firstIndex] > h[secondIndex]
}

func (iHeap *MaxHeap) Swap(firstIndex, secondIndex int) {
	h := *iHeap
	if firstIndex < 0 || secondIndex < 0 || firstIndex >= len(h) || secondIndex >= len(h) {
		panic("index out of range in swap")
	}

	h[firstIndex], h[secondIndex] = h[secondIndex], h[firstIndex]
}

func (iHeap *MaxHeap) Push(x any) {
	value, ok := x.(int)
	if !ok {
		panic("heap: Push of non-int value")
	}

	*iHeap = append(*iHeap, value)
}

func (iHeap *MaxHeap) Pop() any {
	h := *iHeap
	length := len(h)

	if length == 0 {
		return nil
	}

	last := h[length-1]
	*iHeap = h[:length-1]

	return last
}

func readInt(prompt string, min, max int) (int, error) {
	var value int
	_, err := fmt.Scan(&value)

	if err != nil {
		return 0, fmt.Errorf("invalid input of %s: %w", prompt, err)
	}

	if value < min || value > max {
		return 0, fmt.Errorf("%s must be between %d and %d", prompt, min, max)
	}

	return value, nil
}

func main() {
	countOfDishes, err := readInt("count of dishes", 1, 10000)
	if err != nil {
		fmt.Println(err)
		return
	}

	dishHeap := &MaxHeap{}
	heap.Init(dishHeap)

	for range countOfDishes {
		rating, err := readInt("rating of dish", -10000, 10000)
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
