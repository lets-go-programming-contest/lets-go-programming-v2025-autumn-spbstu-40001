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
	NumCode  int     `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    float64 `json:"value"     xml:"Value"`
}

func (v *Valute) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var valueStr string

	err := decoder.DecodeElement(&valueStr, &start)
	if err != nil {
		panic(err)
	}

	normalized := strings.Replace(valueStr, ",", ".", 1)
	v.Value, err = strconv.ParseFloat(normalized, 64)
	if err != nil {
		panic(err)
	}

	return nil
}
