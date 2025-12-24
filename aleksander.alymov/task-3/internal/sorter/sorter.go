package sorter

import "sort"

type Sortable interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

type Sorter interface {
	Sort(data Sortable)
}

type DescendingSorter struct{}

func NewDescendingSorter() Sorter {
	return &DescendingSorter{}
}

func (s *DescendingSorter) Sort(data Sortable) {
	sort.Sort(sort.Reverse(data))
}
