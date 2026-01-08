package currency

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ConvertValues(bank *Bank) (outputList, error) {
	if bank == nil {
		return nil, fmt.Errorf("bank is nil")
	}

	result := make(outputList, 0, len(bank.Items))

	for _, item := range bank.Items {
		strVal := strings.TrimSpace(item.Value)
		strVal = strings.Replace(strVal, ",", ".", 1)

		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return nil, fmt.Errorf("parse value %q: %w", item.Value, err)
		}

		result = append(result, outputItem{
			NumCode:  item.NumCode,
			CharCode: item.CharCode,
			Value:    floatVal,
		})
	}

	return result, nil
}

func (list outputList) SortDesc() {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Value > list[j].Value
	})
}
