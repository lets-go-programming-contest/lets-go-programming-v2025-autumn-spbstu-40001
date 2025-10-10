package main

import (
	"container/heap"
	"fmt"

	"github.com/Aapng-cmd/task-2-2/internal/myheap"
)

func main() {
	var (
		nDishesNumber, actualDish, wanting int
		workHeap                           = &myheap.Heap{}
	)

	_, err := fmt.Scan(&nDishesNumber)
	if err != nil {
		fmt.Println("Invalid dishes count")

		return
	}

	heap.Init(workHeap)

	for range nDishesNumber {
		_, err = fmt.Scan(&actualDish)
		if err != nil {
			fmt.Println("invalid input for dish number")

			return
		}

		err = heap.Push(workHeap, actualDish)
		if err != nil {
		    fmt.Println(err)
		}
	}

	_, err = fmt.Scan(&wanting)
	if err != nil {
		fmt.Println("invalid input for wanted dish")

		return
	}

	for range wanting - 1 {
		if workHeap.Len() == 0 {
			fmt.Println("no dish for you")
			return
		} // здесь надо проверять, потому что Pop может удалить слишком много

		heap.Pop(workHeap)
	}

	fmt.Println(heap.Pop(workHeap))
}
