package xml

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type currencyValue float64

func (cv *currencyValue) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var valueString string

	err := decoder.DecodeElement(&valueString, &start)
	if err != nil {
		return fmt.Errorf("decode element: %w", err)
	}

	valueString = strings.Replace(valueString, ",", ".", 1)

	value, err := strconv.ParseFloat(valueString, 64)
	if err != nil {
		return fmt.Errorf("parse float %q: %w", valueString, err)
	}

	*cv = currencyValue(value)

	return nil
}
