package currencyrate

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type floatWithDots float64

type CurrencyRate struct {
	Valute []*singleValute
}

type singleValute struct {
	NumCode  int           `json:"num_code"  xml:"NumCode"`
	CharCode string        `json:"char_code" xml:"CharCode"`
	Value    floatWithDots `json:"value"     xml:"Value"`
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
