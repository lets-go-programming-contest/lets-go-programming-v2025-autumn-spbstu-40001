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

func (dh *DishHeap) Len() int { return len(*dh) }

func (dh *DishHeap) Less(i, j int) bool {
	return (*dh)[i].rating > (*dh)[j].rating
}

func (dh *DishHeap) Swap(i, j int) {
	(*dh)[i], (*dh)[j] = (*dh)[j], (*dh)[i]
}

func (dh *DishHeap) Push(x any) {
	dish, ok := x.(Dish)
	if !ok {
		panic("invalid data type")
	}
	*dh = append(*dh, dish)
}

func (dh *DishHeap) Pop() any {
	old := *dh
	n := len(old)
	if n == 0 {
		panic("heap is empty")
	}
	x := old[n-1]
	*dh = old[0 : n-1]
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

	dishHeap := &DishHeap{}
	heap.Init(dishHeap)

	for i, rating := range ratings {
		heap.Push(dishHeap, Dish{rating: rating, index: i + 1})
	}

	for i := 0; i < dishCount-1; i++ {
		heap.Pop(dishHeap)
	}

	dish, ok := heap.Pop(dishHeap).(Dish)
	if !ok {
		fmt.Println("Failed to assert type to Dish")
		return
	}

	fmt.Println(dish.index)
}
