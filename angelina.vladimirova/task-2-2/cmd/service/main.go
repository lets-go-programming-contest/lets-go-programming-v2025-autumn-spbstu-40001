package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	if n == 0 {
		return -1
	}
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
		return 0, nil, 0, fmt.Errorf("invalid count: %d", count)
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
		return -1, fmt.Errorf("empty ratings")
	}

	maxHeap := MaxHeap(ratings)
	heap.Init(&maxHeap)

	var result int
	for i := 0; i < positionK; i++ {
		value := heap.Pop(&maxHeap)
		if value == -1 {
			return -1, fmt.Errorf("heap is empty")
		}
		result = value.(int)
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

	fmt.Println(result)
}
