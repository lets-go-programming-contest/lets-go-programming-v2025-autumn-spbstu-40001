package valute

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type (
	ValuteValue float32
	Valute      struct {
		NumCode  int         `json:"num_code"  xml:"NumCode"`
		CharCode string      `json:"char_code" xml:"CharCode"`
		Value    ValuteValue `json:"value"     xml:"Value"`
	}
	ValuteSlice struct {
		XMLName xml.Name `xml:"ValCurs"`
		Valutes []Valute `xml:"Valute"`
	}
)

func (obj *ValuteValue) UnmarshalXML(decode *xml.Decoder, start xml.StartElement) error {
	var value string

	err := decode.DecodeElement(&value, &start)
	if err != nil {
		return fmt.Errorf("failed to decode valute value: %w", err)
	}

	value = strings.ReplaceAll(value, ",", ".")

	temp, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return fmt.Errorf("failed to convert valute value: %w", err)
	}

	*obj = ValuteValue(temp)

	return nil
}
