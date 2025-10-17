package currency

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type (
	ValueWrapper float32
	Currency     struct {
		NumCode  uint         `json:"num_code"  xml:"NumCode"`
		CharCode string       `json:"char_code" xml:"CharCode"`
		Value    ValueWrapper `json:"value"     xml:"Value"`
	}
)

type Rates struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Data    []Currency `xml:"Valute"`
}

func (wrapper *ValueWrapper) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var raw string

	err := decoder.DecodeElement(&raw, &start)
	if err != nil {
		return fmt.Errorf("failed to parse raw Valute: %w", err)
	}

	value, err := strconv.ParseFloat(strings.Replace(raw, ",", ".", 1), 32)
	if err != nil {
		return fmt.Errorf("failed to parse rate value: %w", err)
	}

	*wrapper = ValueWrapper(value)

	return nil
}
