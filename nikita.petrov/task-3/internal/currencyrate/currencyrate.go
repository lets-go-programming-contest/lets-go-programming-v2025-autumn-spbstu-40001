package currencyrate

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type floatWithDots float64

type CurrencyRate struct {
	Valute []struct {
		NumCode  int           `xml:"NumCode"  json:"num_code"`
		CharCode string        `xml:"CharCode" json:"char_code"`
		Value    floatWithDots `xml:"Value"    json:"value"`
	}
}

func (fd *floatWithDots) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	err := d.DecodeElement(&content, &start)

	if err != nil {
		panic(err)
	}

	content = strings.ReplaceAll(content, ",", ".")

	var retFloat64 float64

	retFloat64, err = strconv.ParseFloat(content, 64)
	if err != nil {
		panic(err)
	}

	*fd = floatWithDots(retFloat64)

	return nil
}
