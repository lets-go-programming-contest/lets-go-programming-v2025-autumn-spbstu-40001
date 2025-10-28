package currency

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Rates struct {
	Data []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode  int         `json:"num_code"  xml:"NumCode"`
	CharCode string      `json:"char_code" xml:"CharCode"`
	Value    FloatforCur `json:"value"     xml:"Value"`
}

type FloatforCur float64

func (cf *FloatforCur) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var valueStr string

	err := decoder.DecodeElement(&valueStr, &start)
	if err != nil {
		return fmt.Errorf("failed to parse value: %w", err)
	}

	normalized := strings.Replace(valueStr, ",", ".", 1)

	value, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("invalid number %q: %w", valueStr, err)
	}

	*cf = FloatforCur(value)

	return nil
}
