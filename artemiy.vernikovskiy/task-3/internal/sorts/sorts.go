package sorts

import (
	"sort"

	"github.com/Aapng-cmd/task-3/internal/models"
)

func SortDataByValue(valCurs models.ValCurs) models.ValCurs {
	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	return valCurs
}
