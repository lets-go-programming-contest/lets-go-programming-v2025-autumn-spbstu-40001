package types

import (
	"encoding/xml"
	"sort"
	"strconv"
	"strings"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	XMLName  xml.Name `xml:"Valute"`
	ID       string   `xml:"ID,attr"`
	NumCode  int      `xml:"NumCode"`
	CharCode string   `xml:"CharCode"`
	Nominal  int      `xml:"Nominal"`
	Name     string   `xml:"Name"`
	Value    string   `xml:"Value"`
}

type CurrencyOutput struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func (v *Valute) ConvertToOutput() (*CurrencyOutput, error) {
	valueStr := strings.Replace(v.Value, ",", ".", -1)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return nil, err
	}

	return &CurrencyOutput{
		NumCode:  v.NumCode,
		CharCode: v.CharCode,
		Value:    value,
	}, nil
}

func (vc *ValCurs) SortByValueDesc() []CurrencyOutput {
	var currencies []CurrencyOutput

	for _, valute := range vc.Valutes {
		output, err := valute.ConvertToOutput()
		if err != nil {
			continue
		}
		currencies = append(currencies, *output)
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return currencies
}
