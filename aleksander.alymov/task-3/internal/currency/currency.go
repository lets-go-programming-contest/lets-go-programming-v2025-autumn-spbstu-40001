package currency

import (
	"fmt"
	"sort"

	"github.com/netwite/task-3/internal/json"
	"github.com/netwite/task-3/internal/xml"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type byValueDesc []Currency

func (a byValueDesc) Len() int           { return len(a) }
func (a byValueDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byValueDesc) Less(i, j int) bool { return a[i].Value > a[j].Value }

func ProcessValutes(inputFile, outputFile string) error {
	valutes, err := xml.ParseXML(inputFile)
	if err != nil {
		return fmt.Errorf("parse XML: %w", err)
	}

	currencies := make([]Currency, 0, len(valutes))
	for _, valute := range valutes {
		currency := Currency{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    valute.Value,
		}
		currencies = append(currencies, currency)
	}

	sort.Sort(byValueDesc(currencies))

	if err := json.WriteJSON(currencies, outputFile); err != nil {
		return fmt.Errorf("write JSON: %w", err)
	}

	return nil
}
