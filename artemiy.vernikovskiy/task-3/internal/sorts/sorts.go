// Package sorts provides sorting utilities for currency data.
package sorts

import (
	"sort"

	"github.com/Aapng-cmd/task-3/internal/models"
)

// SortDataByValue sorts the ValCurs data in descending order based on the Value field of each Valute.
// It modifies the input ValCurs in place and returns the sorted ValCurs.
func SortDataByValue(valCurs models.ValCurs) models.ValCurs {
	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	return valCurs
}
