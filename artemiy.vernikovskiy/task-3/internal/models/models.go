package models

import (
	"encoding/xml"
)

type Settings struct {
	InputFileSetting  string `yaml:"input-file"`
	OutputFileSetting string `yaml:"output-file"`
}

type ActualData struct {
	NumCode  string `json:"num_code"  xml:"NumCode"`
	CharCode string `json:"char_code" xml:"CharCode"`
	Value    string `json:"value"     xml:"Value"`
}

type ValCurs struct {
	XMLName xml.Name     `xml:"ValCurs"`
	Date    string       `xml:"Date,attr"`
	Name    string       `xml:"name,attr"`
	Valutes []ActualData `xml:"Valute"`
}
