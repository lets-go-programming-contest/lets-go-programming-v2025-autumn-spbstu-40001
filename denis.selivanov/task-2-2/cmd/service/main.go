package main

import (
	"container/heap"
	"fmt"
)

type Dish struct {
	rating int
	index  int
}

type DishHeap []Dish

func (h DishHeap) Len() int { return len(h) }

func (h DishHeap) Less(i, j int) bool {
	return h[i].rating > h[j].rating
}

func (h DishHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *DishHeap) Push(x interface{}) {
	*h = append(*h, x.(Dish))
}

func (h *DishHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var n, k int
	fmt.Scanf("%d", &n)

	ratings := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ratings[i])
	}

	fmt.Scanf("%d", &k)

	h := &DishHeap{}
	heap.Init(h)

	for i, r := range ratings {
		heap.Push(h, Dish{rating: r, index: i + 1})
	}

	var result int
	for i := 0; i < k; i++ {
		d := heap.Pop(h).(Dish)
		result = d.index
	}

	fmt.Println(result)
}
