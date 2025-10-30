package currency

import (
	"github.com/vikaglushkova/task-3/internal/xml"
	"sort"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ConvertAndSort(valCurs *xml.ValCurs) []Currency {
	currencies := make([]Currency, len(valCurs.Valutes))

	for i, valute := range valCurs.Valutes {
		currencies[i] = Currency{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    valute.Value,
		}
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return currencies
}
