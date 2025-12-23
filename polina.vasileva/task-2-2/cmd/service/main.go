package main

import (
	"container/heap"
	"fmt"

	"polina.vasileva/task-2-2/pkg/intheap"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Panic recovered:", r)
		}
	}()

	var dishNum int

	_, err := fmt.Scan(&dishNum)
	if err != nil {
		fmt.Println("Error reading number of dishes:", err)

		return
	}

	intheap := &intheap.IntHeap{}
	heap.Init(intheap)

	for range dishNum {
		var temp int

		_, err := fmt.Scan(&temp)
		if err != nil {
			fmt.Println("Error reading dish value", err)

			return
		}

		heap.Push(intheap, temp)
	}

	var rating int

	_, err = fmt.Scan(&rating)
	if err != nil {
		fmt.Println("Error reading desired rating position:", err)

		return
	}

	if rating > intheap.Len() {
		fmt.Println("There is no such dish")

		return
	}

	for range rating - 1 {
		heap.Pop(intheap)
	}

	fmt.Println(heap.Pop(intheap))
}
