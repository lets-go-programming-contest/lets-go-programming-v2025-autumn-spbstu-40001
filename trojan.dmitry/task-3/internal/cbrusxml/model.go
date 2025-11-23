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
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    FloatComma `json:"value"     xml:"Value"`
}
