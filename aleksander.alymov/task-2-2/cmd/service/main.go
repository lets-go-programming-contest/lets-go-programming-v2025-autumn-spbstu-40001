package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	value, ok := x.(int)
	if !ok {
		return
	}

	*h = append(*h, value)
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

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
		return -1
	}

	tempHeap := &IntHeap{}
	heap.Init(tempHeap)

	var result int

	index := 0
	for index < kth {
		dish := heap.Pop(k.dishes)

		dishValue, ok := dish.(int)
		if !ok {
			return -1
		}

		if index == kth-1 {
			result = dishValue
		}

		heap.Push(tempHeap, dishValue)

		index++
	}

	for tempHeap.Len() > 0 {
		dish := heap.Pop(tempHeap)

		dishValue, ok := dish.(int)
		if !ok {
			return -1
		}

		heap.Push(k.dishes, dishValue)
	}

	return result
}

func main() {
	var dishCount int
	var preferenceOrder int

	_, err := fmt.Scan(&dishCount)
	if err != nil {
		return
	}

	finder := NewKthPreferenceFinder()

	index := 0
	for index < dishCount {
		var rating int

		_, err := fmt.Scan(&rating)
		if err != nil {
			return
		}

		finder.AddDish(rating)

		index++
	}

	_, err = fmt.Scan(&preferenceOrder)
	if err != nil {
		return
	}

	result := finder.FindKthPreference(preferenceOrder)
	fmt.Println(result)
}
