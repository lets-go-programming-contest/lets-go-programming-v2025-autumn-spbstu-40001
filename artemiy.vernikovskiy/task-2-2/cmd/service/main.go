package main

import (
	"container/heap"
	"fmt"

	"github.com/Aapng-cmd/task-2-2/internal/my_heap"
)

func main() {
	var (
		n, a, wanting int
		workHeap      = &my_heap.Heap{}
	)

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid dishes count")

		return
	}

	heap.Init(workHeap)

	for range n {
		_, err = fmt.Scan(&a)
		if err != nil {
			fmt.Println("invalid input")

			return
		}

		heap.Push(workHeap, a)
	}

	_, err = fmt.Scan(&wanting)
	if err != nil {
		fmt.Println("invalid input")

		return
	}

	for range wanting - 1 {
		if workHeap.Len() == 0 {
			fmt.Println("no dish for you")
		}

		heap.Pop(workHeap)
	}

	fmt.Println(heap.Pop(workHeap))
}
