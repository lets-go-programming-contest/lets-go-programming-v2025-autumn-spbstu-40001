package valcurs

import (
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Currency struct {
	Currencies []struct {
		NumCode  int          `json:"num_code"  xml:"NumCode"`
		CharCode string       `json:"char_code" xml:"CharCode"`
		Value    ParseFloat64 `json:"value"     xml:"Value"`
	} `xml:"Valute"`
}

type ParseFloat64 float64

func (c *ParseFloat64) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var valueStr string

	err := decoder.DecodeElement(&valueStr, &start)
	if err != nil {
		return fmt.Errorf("failed to parse value: %w", err)
	}

	valueStr = strings.Replace(valueStr, ",", ".", 1)

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("failed to parse value '%s': %w", valueStr, err)
	}

	*c = ParseFloat64(value)

	return nil
}

func (v *Currency) SortCurrenciesByValueDesc() {
	sort.Slice(v.Currencies, func(i, j int) bool {
		return v.Currencies[i].Value > v.Currencies[j].Value
	})
}
