package main

import (
	"container/heap"
	"errors"
	"fmt"
)

var (
	ErrInvalidRating     = errors.New("rating must be between -10000 and 10000")
	ErrInvalidPreference = errors.New("preference order out of range")
)

type DishRatingHeap struct {
	ratings []int
}

func (h *DishRatingHeap) Len() int {
	return len(h.ratings)
}

func (h *DishRatingHeap) Less(i, j int) bool {
	return h.ratings[i] > h.ratings[j]
}

func (h *DishRatingHeap) Swap(i, j int) {
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

func readRatings(totalDishes int) (*DishRatingHeap, error) {
	dishRatings := &DishRatingHeap{
		ratings: []int{},
	}
	heap.Init(dishRatings)

	for range totalDishes {
		var currentRating int
		_, err := fmt.Scan(&currentRating)
		if err != nil {
			return nil, fmt.Errorf("reading dish rating: %w", err)
		}

		if currentRating < -10000 || currentRating > 10000 {
			return nil, ErrInvalidRating
		}

		heap.Push(dishRatings, currentRating)
	}

	return dishRatings, nil
}

func readPreferenceOrder(totalDishes int) (int, error) {
	var preferenceOrder int

	_, err := fmt.Scan(&preferenceOrder)
	if err != nil {
		return 0, fmt.Errorf("reading preference order: %w", err)
	}

	if preferenceOrder < 1 || preferenceOrder > totalDishes {
		return 0, fmt.Errorf("%w: must be between 1 and %d", ErrInvalidPreference, totalDishes)
	}

	return preferenceOrder, nil
}

func findKthPreference(dishRatings *DishRatingHeap, preferenceOrder int) (int, error) {
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
			return 0, errors.New("not enough elements in heap")
		}

		rating, ok := item.(int)
		if !ok {
			return 0, errors.New("unexpected type in heap")
		}

		result = rating
	}

	return result, nil
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

	dishRatings, err := readRatings(totalDishes)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	preferenceOrder, err := readPreferenceOrder(totalDishes)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	result, err := findKthPreference(dishRatings, preferenceOrder)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println(result)
}
