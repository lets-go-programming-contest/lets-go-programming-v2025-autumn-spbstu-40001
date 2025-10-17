package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int) {
	if i >= 0 && i < len(h) && j >= 0 && j < len(h) {
		h[i], h[j] = h[j], h[i]
	}
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	if len(*h) == 0 {
		return nil
	}
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKLargest(foodRatings []int, prefferedDihes int) int {
	if prefferedDihes <= 0 || prefferedDihes > len(foodRatings) {
		panic("Going beyond the boundaries of the array")
	}

	h := &MinHeap{}
	heap.Init(h)

	for i := 0; i < prefferedDihes; i++ {
		heap.Push(h, foodRatings[i])
	}

	for i := prefferedDihes; i < len(foodRatings); i++ {
		if foodRatings[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, foodRatings[i])
		}
	}

	return (*h)[0]
}

func main() {
	var dishesNumber, prefferedDishes int

	_, err := fmt.Scan(&dishesNumber)
	if dishesNumber <= 0 || err != nil {
		fmt.Println("Incorrect number of dishes: ", err)

		return
	}

	foodRatings := make([]int, dishesNumber)
	for i := 0; i < dishesNumber; i++ {
		_, err := fmt.Scan(&foodRatings[i])
		if err != nil {
			fmt.Printf("Error reading dish %d: %v\n", i+1, err)

			return
		}
	}

	_, err = fmt.Scan(&prefferedDishes)
	if err != nil || prefferedDishes <= 0 || prefferedDishes > dishesNumber {
		fmt.Println("Incorrect preferred dishes number:", err)

		return
	}

	answer := findKLargest(foodRatings, prefferedDishes)
	fmt.Println(answer)
}
