package sorter

import (
	"sort"

	"github.com/Elektrek/task-3/internal/model"
)

func SortByValueDescending(currencies []model.Currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}
