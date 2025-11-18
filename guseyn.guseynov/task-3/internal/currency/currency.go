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
	type valuteAlias struct {
		NumCode  int    `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	}

	var alias valuteAlias

	err := decoder.DecodeElement(&alias, &start)
	if err != nil {
		panic(err)
	}

	v.NumCode = alias.NumCode
	v.CharCode = alias.CharCode

	normalized := strings.Replace(strings.TrimSpace(alias.Value), ",", ".", 1)

	v.Value, err = strconv.ParseFloat(normalized, 64)
	if err != nil {
		panic(err)
	}

	return nil
}
