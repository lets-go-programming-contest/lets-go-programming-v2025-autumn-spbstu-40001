package intheap

import "cmp"

func New(data []int, comparator func(int, int) bool) Heap {
	return Heap{data, comparator}
}

type Heap struct {
	data       []int
	comparator func(int, int) bool
}

func (obj *Heap) Len() int {
	return len(obj.data)
}

func (obj *Heap) Less(lhs, rhs int) bool {
	if lhs >= obj.Len() || rhs >= obj.Len() {
		panic("invalid index range")
	}

	if obj.comparator == nil {
		return cmp.Less(obj.data[lhs], obj.data[rhs])
	}

	return obj.comparator(obj.data[lhs], obj.data[rhs])
}

func (obj *Heap) Swap(lhs, rhs int) {
	if lhs >= obj.Len() || rhs >= obj.Len() {
		panic("invalid index range")
	}

	obj.data[lhs], obj.data[rhs] = obj.data[rhs], obj.data[lhs]
}

func (obj *Heap) Push(data any) {
	iData, ok := data.(int)
	if !ok {
		panic("unexpected data type (expected: int)")
	}

	obj.data = append(obj.data, iData)
}

func (obj *Heap) Pop() any {
	size := obj.Len()
	if size == 0 {
		return nil
	}

	data := obj.data[size-1]
	obj.data = obj.data[0 : size-1]

	return data
}
