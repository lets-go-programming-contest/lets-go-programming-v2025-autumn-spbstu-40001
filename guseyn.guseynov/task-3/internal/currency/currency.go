package currency

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type ValCurs struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Valutes []Valute   `xml:"Valute"`
}

type Valute struct {
	ID       string    `xml:"ID,attr"`
	NumCode  int       `json:"num_code" xml:"NumCode"`
	CharCode string    `json:"char_code" xml:"CharCode"`
	Nominal  int       `xml:"Nominal"`
	Name     string    `xml:"Name"`
	Value    float64   `json:"value" xml:"Value"`
}

func (v *Valute) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	type alias struct {
		ID       string `xml:"ID,attr"`
		NumCode  int    `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Nominal  int    `xml:"Nominal"`
		Name     string `xml:"Name"`
		Value    string `xml:"Value"`
	}

	var a alias
	err := decoder.DecodeElement(&a, &start)
	if err != nil {
		return err
	}

	v.ID = a.ID
	v.NumCode = a.NumCode
	v.CharCode = a.CharCode
	v.Nominal = a.Nominal
	v.Name = a.Name

	normalized := strings.Replace(a.Value, ",", ".", 1)
	v.Value, err = strconv.ParseFloat(normalized, 64)
	if err != nil {
		return err
	}

	return nil
}
