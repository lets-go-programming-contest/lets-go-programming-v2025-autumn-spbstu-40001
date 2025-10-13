package main

import (
	"container/heap"
	"fmt"

	"github.com/A1exCRE/task-2-2/internal/intheap"
)

func main() {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("Error:", rec)
		}
	}()

	var dishCount int

	_, err := fmt.Scan(&dishCount)
	if err != nil {
		fmt.Println("Invalid input", err)
		return
	}

	ratingsHeap := &intheap.IntHeap{}
	heap.Init(ratingsHeap)

	for range dishCount {
		var rating int

		_, err := fmt.Scan(&rating)
		if err != nil {
			fmt.Println("Invalid input", err)
			return
		}

		heap.Push(ratingsHeap, rating)
	}

	var preferredDishNumber int

	_, err = fmt.Scan(&preferredDishNumber)
	if err != nil {
		fmt.Println("Invalid input", err)
		return
	}

	for range preferredDishNumber - 1 {
		heap.Pop(ratingsHeap)
	}

	result := heap.Pop(ratingsHeap)
	fmt.Println(result)
}
