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
	*h = append(*h, x.(int))
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

	for i := 0; i < kth; i++ {
		dish := heap.Pop(k.dishes).(int)
		if i == kth-1 {
			result = dish
		}
		heap.Push(tempHeap, dish)
	}

	for tempHeap.Len() > 0 {
		heap.Push(k.dishes, heap.Pop(tempHeap))
	}

	return result
}

func main() {
	var n, k int

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Ошибка ввода количества блюд:", err)
		return
	}

	finder := NewKthPreferenceFinder()

	for i := 0; i < n; i++ {
		var rating int
		_, err := fmt.Scan(&rating)
		if err != nil {
			fmt.Println("Ошибка ввода рейтинга:", err)
			return
		}
		finder.AddDish(rating)
	}

	_, err = fmt.Scan(&k)
	if err != nil {
		fmt.Println("Ошибка ввода k:", err)
		return
	}

	result := finder.FindKthPreference(k)
	fmt.Println(result)
}
