package currency

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    FloatValue `json:"value"     xml:"Value"`
}

type FloatValue float64

func (fv *FloatValue) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var valueStr string

	err := decoder.DecodeElement(&valueStr, &start)
	if err != nil {
		panic(err)
	}

	normalized := strings.Replace(strings.TrimSpace(valueStr), ",", ".", 1)

	value, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		panic(err)
	}

	*fv = FloatValue(value)

	return nil
}
