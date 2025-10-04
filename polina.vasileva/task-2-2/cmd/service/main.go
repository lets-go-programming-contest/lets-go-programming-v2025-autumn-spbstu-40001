package main

import (
	"container/heap"
	"fmt"

	"polina.vasileva/task-2-2/pkg/intheap"
)

func main() {
	var dishNum int

	_, err := fmt.Scan(&dishNum)
	if err != nil {
		fmt.Println("Invalid input", err)

		return
	}

	intheap := &intheap.IntHeap{}
	heap.Init(intheap)

	for range dishNum {
		var temp int

		_, err := fmt.Scan(&temp)
		if err != nil {
			fmt.Println("Invalid input", err)

			return
		}

		heap.Push(intheap, temp)
	}

	var rating int

	_, err = fmt.Scan(&rating)
	if err != nil {
		fmt.Println("Invalid input", err)

		return
	}

	for range rating - 1 {
		if intheap.Len() == 0 {
			fmt.Println("There is no such dish")
		}

		heap.Pop(intheap)
	}

	fmt.Println(heap.Pop(intheap))
}
