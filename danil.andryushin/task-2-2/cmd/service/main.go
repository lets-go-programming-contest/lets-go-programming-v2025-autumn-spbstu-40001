package main

import (
	"container/heap"
	"fmt"

	"github.com/atroxxxxxx/task-2-2/internal/task"
)

func Greater(a, b int) bool {
	return a > b
}

func main() {
	test := &task.Heap{Functor: Greater}
	heap.Init(test)
	heap.Push(test, 3)
	heap.Push(test, 5)
	heap.Push(test, 7)
	heap.Push(test, 9)
	heap.Push(test, 2)
	heap.Push(test, 6)
	heap.Push(test, 12)
	fmt.Println(*test)
	heap.Pop(test)
	fmt.Println(heap.Pop(test))
}
