package types

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Rates struct {
	Data []Types `xml:"Valute"`
}

type Types struct {
	NumCode  int         `json:"num_code"  xml:"NumCode"`
	CharCode string      `json:"char_code" xml:"CharCode"`
	Value    FloatTypes `json:"value"     xml:"Value"`
}

type FloatTypes float64

func (cf *FloatTypes) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var valueStr string

	err := decoder.DecodeElement(&valueStr, &start)
	if err != nil {
		return fmt.Errorf("failed to parse value: %w", err)
	}

	fixed := strings.Replace(valueStr, ",", ".", 1)

	value, err := strconv.ParseFloat(fixed, 64)
	if err != nil {
		return fmt.Errorf("invalid number %q: %w", valueStr, err)
	}

	*cf = FloatTypes(value)

	return nil
}
