package currency

import (
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Currency struct {
	NumCode  int     `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    float64 `json:"value"     xml:"Value"`
}

func (c *Currency) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	type alias Currency

	var temp struct {
		alias
		Value string `xml:"Value"`
	}

	err := decoder.DecodeElement(&temp, &start)
	if err != nil {
		return fmt.Errorf("decode element: %w", err)
	}

	*c = Currency(temp.alias)

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
