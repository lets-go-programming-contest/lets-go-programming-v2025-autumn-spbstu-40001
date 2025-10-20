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
		fmt.Println("Failed to read number of dishes:", err)

		return
	}

	ratingsHeap := &intheap.IntHeap{}
	heap.Init(ratingsHeap)

	for range dishCount {
		var rating int

		_, err := fmt.Scan(&rating)
		if err != nil {
			fmt.Println("Failed to read dish rating:", err)

			return
		}

		heap.Push(ratingsHeap, rating)
	}

	var preferredDishNumber int

	_, err = fmt.Scan(&preferredDishNumber)
	if err != nil {
		fmt.Println("Failed to read preferred dish number:", err)

		return
	}

	for range preferredDishNumber - 1 {
		heap.Pop(ratingsHeap)
	}

	result := heap.Pop(ratingsHeap)
	if result == nil {
		fmt.Println("No result found:", err)

		return
	}

	fmt.Println(result)
}
