package main

import (
	"container/heap"
	"fmt"

	"github.com/A1exCRE/task-2-2/internal/intheap"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
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
			fmt.Println("Invalid input")
			return
		}

		heap.Push(ratingsHeap, rating)
	}

	fmt.Println("Куча:", *ratingsHeap)
}
