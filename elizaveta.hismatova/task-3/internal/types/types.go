package types

import (
	"encoding/xml"
	"fmt"
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

func (v Valute) ToOutput() (CurrencyOutput, error) {
	cleanedValue := strings.ReplaceAll(v.Value, ",", ".")
	value, err := strconv.ParseFloat(cleanedValue, 64)

	if err != nil {
		return CurrencyOutput{}, fmt.Errorf("failed to parse value: %w", err)
	}

	return CurrencyOutput{
		NumCode:  v.NumCode,
		CharCode: v.CharCode,
		Value:    value,
	}, nil
}

func (vc ValCurs) SortByValueDesc() []CurrencyOutput {
	outputs := make([]CurrencyOutput, 0, len(vc.Valutes))

	for _, valute := range vc.Valutes {
		output, err := valute.ToOutput()

		if err != nil {
			continue
		}

		outputs = append(outputs, output)
	}

	sort.Slice(outputs, func(i, j int) bool {
		return outputs[i].Value > outputs[j].Value
	})

	return outputs
}
