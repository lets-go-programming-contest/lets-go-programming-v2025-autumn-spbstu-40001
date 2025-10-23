package models

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type CommaFloat float64 // please, let this live

func (cf *CommaFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// let it be here, please, i do not want one more packet in this small task
	var sIsAWorkingStringForFloatsWithCommaIsThisNameLongEnough string

	err := d.DecodeElement(&sIsAWorkingStringForFloatsWithCommaIsThisNameLongEnough, &start)
	if err != nil {
		return fmt.Errorf("ah, kozache, UnmarshalXML override func failed: %w", err)
	}

	sIsAWorkingStringForFloatsWithCommaIsThisNameLongEnough = strings.ReplaceAll(
		sIsAWorkingStringForFloatsWithCommaIsThisNameLongEnough,
		",",
		".",
	)

	val, err := strconv.ParseFloat(sIsAWorkingStringForFloatsWithCommaIsThisNameLongEnough, 64)
	if err != nil {
		return fmt.Errorf("ah, kozache, UnmarshalXML override func failed: %w", err)
	}

	*cf = CommaFloat(val)

	return nil
}

type Settings struct {
	InputFileSetting  string `yaml:"input-file"`
	OutputFileSetting string `yaml:"output-file"`
}

type ActualData struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    CommaFloat `json:"value"     xml:"Value"`
}

type ValCurs struct {
	Valutes []ActualData `xml:"Valute"`
}
