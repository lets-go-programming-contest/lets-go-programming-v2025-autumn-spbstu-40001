package cbrusxml

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type FloatComma float64

func (f *FloatComma) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string

	if err := d.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode element: %w", err)
	}

	raw = strings.TrimSpace(raw)
	raw = strings.ReplaceAll(raw, ",", ".")

	if raw == "" {
		*f = 0
		return nil
	}

	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return fmt.Errorf("parse float: %w", err)
	}

	*f = FloatComma(val)
	return nil
}

type Valute struct {
	NumCode  int        `xml:"NumCode"  json:"num_code"`
	CharCode string     `xml:"CharCode" json:"char_code"`
	Value    FloatComma `xml:"Value"    json:"value"`
}
