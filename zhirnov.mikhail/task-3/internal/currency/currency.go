package currency

import (
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type ValCurs struct {
	Valutes []struct {
		NumCode  int           `json:"num_code"  xml:"NumCode"`
		CharCode string        `json:"char_code" xml:"CharCode"`
		Value    CustomFloat64 `json:"value"     xml:"Value"`
	} `xml:"Valute"`
}

type CustomFloat64 float64

func (c *CustomFloat64) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var valueStr string

	err := decoder.DecodeElement(&valueStr, &start)
	if err != nil {
		return fmt.Errorf("failed to parse value: %w", err)
	}

	value, err := strconv.ParseFloat(strings.Replace(valueStr, ",", ".", 1), 64)
	if err != nil {
		return fmt.Errorf("failed to parse value: %w", err)
	}

	*c = CustomFloat64(value)

	return nil
}

func (v *ValCurs) SortByValueDesc() {
	sort.Slice(v.Valutes, func(i, j int) bool {
		return v.Valutes[i].Value > v.Valutes[j].Value
	})
}
