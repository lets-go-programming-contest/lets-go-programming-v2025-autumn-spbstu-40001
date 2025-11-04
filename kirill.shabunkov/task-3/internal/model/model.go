package model

import (
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int     `json:"num_code" xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    float64 `json:"value" xml:"Value"`
}

func (v *Valute) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type tempStruct struct {
		NumCode  int    `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	}

	var t tempStruct

	err := d.DecodeElement(&t, &start)
	if err != nil {
		return fmt.Errorf("сannot decode XML element: %w", err)
	}

	strValue := strings.Replace(t.Value, ",", ".", -1)
	value, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		return fmt.Errorf("cannot parse value '%s': %w", t.Value, err)
	}

	v.NumCode = t.NumCode
	v.CharCode = t.CharCode
	v.Value = value

	return nil
}

func (v *ValCurs) SortByValue() {
	sort.Slice(v.Valutes, func(i, j int) bool {
		return v.Valutes[i].Value > v.Valutes[j].Value
	})
}
