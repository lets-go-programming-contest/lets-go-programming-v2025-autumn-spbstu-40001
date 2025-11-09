package model

import "sort"

type CurrencyItem struct {
	NumericCode int     `json:"num_code"`
	Code        string  `json:"char_code"`
	Rate        float64 `json:"value"`
}

type CurrencyList struct {
	Items []CurrencyItem
}

func (cl *CurrencyList) OrderByValue() {
	sort.Slice(cl.Items, func(i, j int) bool {
		return cl.Items[i].Rate > cl.Items[j].Rate
	})
}
