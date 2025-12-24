package currency

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type RateValue float64

type ValCurs struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Valutes []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode  int       `json:"num_code"  xml:"NumCode"`
	CharCode string    `json:"char_code" xml:"CharCode"`
	Value    RateValue `json:"value"     xml:"Value"`
}

func (rate *RateValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return fmt.Errorf("failed to decode XML element: %w", err)
	}

	str = strings.Replace(str, ",", ".", 1)

	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("failed to parse rate value '%s': %w", str, err)
	}

	*rate = RateValue(val)

	return nil
}
