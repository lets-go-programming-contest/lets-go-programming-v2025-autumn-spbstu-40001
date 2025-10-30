package primheap

import (
	"cmp"
	"container/heap"
	"errors"
)

var (
	ErrEmptyHeapTop     = errors.New("cannot get top element from empty heap")
	ErrHeapUnderflow    = errors.New("cannot pop from empty heap")
	ErrPopNInvalidCount = errors.New("cannot pop less then 1 or more then Len elements")
	ErrTypecastFailed   = errors.New("cannot convert PrimHeap::inner_stack::pop() return value")
)

type PrimHeap[T cmp.Ordered] struct {
	stack extendedStack[T]
}

func (obj *PrimHeap[T]) Init() {
	heap.Init(&obj.stack)
}

func (obj *PrimHeap[T]) Len() int {
	return obj.stack.Len()
}

func (obj *PrimHeap[T]) Top() (T, error) {
	if obj.stack.Len() == 0 {
		var result T

		return result, ErrEmptyHeapTop
	}

	return obj.stack.data[0], nil
}

func (obj *PrimHeap[T]) Push(value T) {
	heap.Push(&obj.stack, value)
}

func (obj *PrimHeap[T]) Pop() (T, error) {
	result := heap.Pop(&obj.stack)
	if result == nil {
		var result T

		return result, ErrHeapUnderflow
	}

	castedResult, ok := result.(T)
	if !ok {
		return castedResult, ErrTypecastFailed
	}

	return castedResult, nil
}

func (obj *PrimHeap[T]) PopN(count int) (T, error) {
	var result T
	if (count < 1) || (count > obj.Len()) {
		return result, ErrPopNInvalidCount
	}

	for range count - 1 {
		_, err := obj.Pop()
		if err != nil {
			return result, err
		}
	}

	return obj.Pop()
}

func New[T cmp.Ordered](less func(T, T) bool, values ...T) PrimHeap[T] {
	result := PrimHeap[T]{extendedStack[T]{values, less}}
	result.Init()

	return result
}
