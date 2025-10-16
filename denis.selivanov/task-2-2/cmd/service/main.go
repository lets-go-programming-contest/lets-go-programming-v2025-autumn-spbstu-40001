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

func (h *DishHeap) Len() int { return len(*h) }

func (h *DishHeap) Less(i, j int) bool {
	return (*h)[i].rating > (*h)[j].rating
}

func (h *DishHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *DishHeap) Push(x any) {
	dish, ok := x.(Dish)
	if !ok {
		panic("invalid data type")
	}
	*h = append(*h, dish)
}

func (h *DishHeap) Pop() any {
	old := *h
	n := len(old)
	if n == 0 {
		panic("heap is empty")
	}
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic recovered:", err)
		}
	}()

	var numberDishes, dishCount int
	_, err := fmt.Scan(&numberDishes)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	if numberDishes <= 0 {
		fmt.Println("Number of dishes must be greater than 0.")
		return
	}

	ratings := make([]int, numberDishes)
	for i := 0; i < numberDishes; i++ {
		_, err = fmt.Scan(&ratings[i])
		if err != nil {
			fmt.Println("Invalid input:", err)
			return
		}
	}

	_, err = fmt.Scan(&dishCount)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	if dishCount <= 0 || dishCount > numberDishes {
		fmt.Println("Dish count must be between 1 and", numberDishes)
		return
	}

	h := &DishHeap{}
	heap.Init(h)

	for i, rating := range ratings {
		heap.Push(h, Dish{rating: rating, index: i + 1})
	}

	for i := 0; i < dishCount-1; i++ {
		heap.Pop(h)
	}

	dish := heap.Pop(h).(Dish)
	fmt.Println(dish.index)
}
