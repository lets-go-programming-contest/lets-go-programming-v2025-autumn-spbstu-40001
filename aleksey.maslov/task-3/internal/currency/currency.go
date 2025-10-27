package currency

import (
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type CurrencyValue float64

type Valute struct {
	NumCode  int           `json:"num_code" xml:"NumCode"`
	CharCode string        `json:"char_code" xml:"CharCode"`
	Value    CurrencyValue `json:"value" xml:"Value"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

func (cv *CurrencyValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string

	err := d.DecodeElement(&str, &start)
	if err != nil {
		return fmt.Errorf("failed to decode currency value: %w", err)
	}

	str = strings.ReplaceAll(str, ",", ".")

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to parse currency value: %w", err)
	}

	*cv = CurrencyValue(value)

	return nil
}

func SortByValue(valutes []Valute) {
	sort.Slice(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})
}
