package main

import (
	"container/heap"
	"errors"
	"fmt"
)

type MaxHeap []int

var (
	ErrInvalidPrefferedDishes = errors.New("invalid preffered dishec count")
	ErrUnexpectedTypeFromHeap = errors.New("unexpected type from heap")
)

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	if i < 0 || i >= h.Len() || j < 0 || j >= h.Len() {
		panic("Index out of range")
	}

	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	if i < 0 || i >= h.Len() || j < 0 || j >= h.Len() {
		panic("Index out of range")
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x any) {
	num, ok := x.(int)
	if !ok {
		panic("type assertion to int failed")
	}

	*h = append(*h, num)
}

func (h *MaxHeap) Pop() any {
	if len(*h) == 0 {
		return nil
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func findKLargest(foodRatings []int, prefferedDishes int) (int, error) {
	if prefferedDishes <= 0 || prefferedDishes > len(foodRatings) {
		return 0, ErrInvalidPrefferedDishes
	}

	maxHeap := MaxHeap(foodRatings)
	heap.Init(&maxHeap)

	for range prefferedDishes - 1 {
		heap.Pop(&maxHeap)
	}

	item := heap.Pop(&maxHeap)
	num, ok := item.(int)
	if !ok {
		return 0, ErrUnexpectedTypeFromHeap
	}

	return num, nil
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

	answer, err := findKLargest(foodRatings, prefferedDishes)
	if err != nil {
		fmt.Printf("Failed to find %d-th largest rating: %v\n", prefferedDishes, err)

		return
	}

	fmt.Println(answer)
}
