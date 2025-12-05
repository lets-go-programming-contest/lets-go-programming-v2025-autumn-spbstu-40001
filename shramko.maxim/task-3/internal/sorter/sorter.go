package sorter

import (
	"sort"

	"github.com/Elektrek/task-3/internal/model"
)

func SortByValueDescending(currencyItems []model.Currency) {
	sort.Slice(currencyItems, func(i, j int) bool {
		return currencyItems[i].Value > currencyItems[j].Value
	})
}
