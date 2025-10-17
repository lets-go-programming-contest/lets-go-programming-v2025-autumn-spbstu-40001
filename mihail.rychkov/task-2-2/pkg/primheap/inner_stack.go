package primheap

import "cmp"

type extendedStack[T cmp.Ordered] struct {
	data       []T
	comparator func(T, T) bool
}

func (obj *extendedStack[T]) Less(lhsIdx, rhsIdx int) bool {
	obj.validateIndex(lhsIdx)
	obj.validateIndex(rhsIdx)

	if obj.comparator != nil {
		return obj.comparator(obj.data[lhsIdx], obj.data[rhsIdx])
	}

	return cmp.Less(obj.data[lhsIdx], obj.data[rhsIdx])
}

func (obj *extendedStack[T]) Len() int {
	return len(obj.data)
}

func (obj *extendedStack[T]) Swap(idx1, idx2 int) {
	obj.validateIndex(idx1)
	obj.validateIndex(idx2)
	obj.data[idx1], obj.data[idx2] = obj.data[idx2], obj.data[idx1]
}

func (obj *extendedStack[T]) Push(value any) {
	intValue, ok := value.(T)
	if !ok {
		panic("failed to cast heap.push value to stored type")
	}

	obj.data = append(obj.data, intValue)
}

func (obj *extendedStack[T]) Pop() any {
	if obj.Len() == 0 {
		return nil
	}

	var empty T

	result := obj.data[obj.Len()-1]
	obj.data[obj.Len()-1] = empty
	obj.data = obj.data[0 : obj.Len()-1]

	return result
}

func (obj *extendedStack[T]) validateIndex(i int) {
	if (i < 0) || (i >= obj.Len()) {
		panic("got invalid index")
	}
}
