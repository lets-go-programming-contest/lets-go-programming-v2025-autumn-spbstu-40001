package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *MinHeap) Swap(i, j int) {
	if i >= 0 && i < len(*h) && j >= 0 && j < len(*h) {
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	}
}

func (h *MinHeap) Push(x any) {
	num, ok := x.(int)
	if !ok {
		panic("type assertion to int failed")
	}
	*h = append(*h, num)
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

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, rating := range foodRatings[:prefferedDihes] {
		heap.Push(minHeap, rating)
	}

	for _, rating := range foodRatings[prefferedDihes:] {
		if rating > (*minHeap)[0] {
			heap.Pop(minHeap)
			heap.Push(minHeap, rating)
		}
	}

	return (*minHeap)[0]
}

func main() {
	var dishesNumber, prefferedDishes int

	_, err := fmt.Scan(&dishesNumber)
	if dishesNumber <= 0 || err != nil {
		fmt.Println("Incorrect number of dishes: ", err)

		return
	}

	foodRatings := make([]int, dishesNumber)
	for i := range dishesNumber {
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
