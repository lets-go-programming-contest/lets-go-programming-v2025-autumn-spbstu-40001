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
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	value, ok := x.(int)
	if !ok {
		fmt.Printf("Error: attempt to add element of wrong type to heap\n")

		return
	}

	*h = append(*h, value)
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	length := len(old)

	if length == 0 {
		fmt.Printf("Error: attempt to extract element from empty heap\n")

		return -1
	}

	x := old[length-1]
	*h = old[0 : length-1]

	return x
}

type KthPreferenceFinder struct {
	dishes *IntHeap
}

func NewKthPreferenceFinder() *KthPreferenceFinder {
	h := &IntHeap{}
	heap.Init(h)

	return &KthPreferenceFinder{dishes: h}
}

func (k *KthPreferenceFinder) AddDish(rating int) {
	heap.Push(k.dishes, rating)
}

func (k *KthPreferenceFinder) FindKthPreference(kth int) int {
	if kth < 1 || kth > k.dishes.Len() {
		fmt.Printf("Error: invalid kth value = %d\n", kth)

		return -1
	}

	temp := make([]int, 0, kth)
	var result int

	for i := range kth {
		dish := heap.Pop(k.dishes)
		dishValue, ok := dish.(int)

		if !ok {
			fmt.Printf("Error: received element of wrong type from heap\n")

			return -1
		}

		if i == kth-1 {
			result = dishValue
		}

		temp = append(temp, dishValue)
	}

	for _, value := range temp {
		heap.Push(k.dishes, value)
	}

	return result
}

func main() {
	var dishCount int

	_, err := fmt.Scan(&dishCount)
	if err != nil {
		fmt.Printf("Error reading number of dishes: %v\n", err)

		return
	}

	if dishCount <= 0 {
		fmt.Printf("Error: number of dishes must be positive\n")

		return
	}

	finder := NewKthPreferenceFinder()

	for i := range dishCount {
		var rating int
		_, err := fmt.Scan(&rating)

		if err != nil {
			fmt.Printf("Error reading dish rating: %v\n", err)

			return
		}

		finder.AddDish(rating)
	}

	var preferenceOrder int
	_, err = fmt.Scan(&preferenceOrder)

	if err != nil {
		fmt.Printf("Error reading preference order: %v\n", err)

		return
	}

	result := finder.FindKthPreference(preferenceOrder)
	fmt.Println(result)
}
