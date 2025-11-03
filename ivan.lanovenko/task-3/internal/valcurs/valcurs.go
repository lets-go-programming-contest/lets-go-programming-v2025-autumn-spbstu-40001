package valcurs

import (
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type FloatWithComma float64

type ValCurs struct {
	Valutes []struct {
		NumCode  int            `json:"num_code"  xml:"NumCode"`
		CharCode string         `json:"char_code" xml:"CharCode"`
		Value    FloatWithComma `json:"value"     xml:"Value"`
	} `xml:"Valute"`
}

func (f *FloatWithComma) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	valueStr := ""
	if err := d.DecodeElement(&valueStr, &start); err != nil {
		return fmt.Errorf("decode element: %w", err)
	}

	valueStr = strings.ReplaceAll(strings.TrimSpace(valueStr), ",", ".")

	val, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("parse float %q: %w", valueStr, err)
	}

	*f = FloatWithComma(val)

	return nil
}

func (v *ValCurs) SortByValueDown() {
	sort.Slice(v.Valutes, func(i, j int) bool {
		return v.Valutes[i].Value > v.Valutes[j].Value
	})
}
