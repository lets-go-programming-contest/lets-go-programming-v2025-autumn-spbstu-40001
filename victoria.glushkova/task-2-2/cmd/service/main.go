package main

import (
	"container/heap"
	"fmt"
)

type DishRatingHeap struct {
	ratings []int
}

func (h DishRatingHeap) Len() int {
	return len(h.ratings)
}

func (h DishRatingHeap) Less(i, j int) bool {
	return h.ratings[i] > h.ratings[j]
}

func (h DishRatingHeap) Swap(i, j int) {
	h.ratings[i], h.ratings[j] = h.ratings[j], h.ratings[i]
}

func (h *DishRatingHeap) Push(x interface{}) {
	rating, ok := x.(int)
	if !ok {
		fmt.Println("Error: expected integer value")

		return
	}

	h.ratings = append(h.ratings, rating)
}

func (h *DishRatingHeap) Pop() interface{} {
	if len(h.ratings) == 0 {
		fmt.Println("Error: cannot pop from empty heap")

		return nil
	}

	n := len(h.ratings)
	item := h.ratings[n-1]
	h.ratings = h.ratings[:n-1]

	return item
}

func main() {
	var totalDishes int

	_, err := fmt.Scan(&totalDishes)
	if err != nil {
		fmt.Println("Error reading number of dishes:", err)

		return
	}

	if totalDishes < 1 || totalDishes > 10000 {
		fmt.Println("Error: number of dishes must be between 1 and 10000")

		return
	}

	dishRatings := &DishRatingHeap{
		ratings: []int{},
	}
	heap.Init(dishRatings)

	for range totalDishes {
		var currentRating int
		_, err := fmt.Scan(&currentRating)
		if err != nil {
			fmt.Println("Error reading dish rating:", err)

			return
		}

		if currentRating < -10000 || currentRating > 10000 {
			fmt.Println("Error: rating must be between -10000 and 10000")

			return
		}

		heap.Push(dishRatings, currentRating)
	}

	var preferenceOrder int

	_, err = fmt.Scan(&preferenceOrder)
	if err != nil {
		fmt.Println("Error reading preference order:", err)

		return
	}

	if preferenceOrder < 1 || preferenceOrder > totalDishes {
		fmt.Printf("Error: preference order must be between 1 and %d\n", totalDishes)

		return
	}

	tempHeap := &DishRatingHeap{
		ratings: []int{},
	}
	heap.Init(tempHeap)

	for _, rating := range dishRatings.ratings {
		heap.Push(tempHeap, rating)
	}

	var result int
	for range preferenceOrder {
		item := heap.Pop(tempHeap)
		if item == nil {
			fmt.Println("Error: not enough elements in heap")

			return
		}

		rating, ok := item.(int)
		if !ok {
			fmt.Println("Error: unexpected type in heap")

			return
		}

		result = rating
	}

	fmt.Println(result)
}
