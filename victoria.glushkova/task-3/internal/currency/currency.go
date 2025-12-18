package currency

import (
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Currency struct {
	NumCode  int     `json:"num_code" xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    float64 `json:"value" xml:"Value"`
}

func (c *Currency) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type currencyXML struct {
		NumCode  int    `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	}
	var temp currencyXML
	if err := d.DecodeElement(&temp, &start); err != nil {
		return fmt.Errorf("decode element: %w", err)
	}
	c.NumCode = temp.NumCode
	c.CharCode = temp.CharCode
	valueString := strings.Replace(temp.Value, ",", ".", 1)
	value, err := strconv.ParseFloat(valueString, 64)
	if err != nil {
		return fmt.Errorf("parse float %q: %w", valueString, err)
	}
	c.Value = value
	return nil
}

func ConvertAndSort(currencies []Currency) []Currency {
	result := make([]Currency, len(currencies))
	copy(result, currencies)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})
	return result
}
