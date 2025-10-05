package main

import (
	"container/heap"
	"fmt"

	"github.com/atroxxxxxx/task-2-2/internal/extended_stack"
)

func Greater(a, b int) bool {
	return a > b
}

func main() {
	var (
		nDishes int
		dishes  = &extended_stack.Stack{Functor: Greater}
	)
	heap.Init(dishes)
	_, err := fmt.Scan(&nDishes)
	if err != nil || nDishes < 0 {
		fmt.Println("invalid dishes count:")
		return
	}
	var dishRating int
	for range nDishes {
		_, err = fmt.Scan(&dishRating)
		if err != nil {
			fmt.Println("wrong dish rating format:", err)
			return
		}
		heap.Push(dishes, dishRating)
	}
	var dishNumber int
	_, err = fmt.Scan(&dishNumber)
	if err != nil || dishNumber < 0 || dishNumber > nDishes {
		fmt.Println("invalid dish number")
	}
	for range dishNumber - 1 {
		heap.Pop(dishes)
	}
	fmt.Println(heap.Pop(dishes))
}
