// Package models defines data structures for currency exchange rates and configuration settings.
package models

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

// CommaFloat represents a float64 value that is unmarshaled from
// XML strings containing commas as decimal separators.
type CommaFloat float64

// UnmarshalXML implements xml.Unmarshaler for CommaFloat.
// It converts comma-separated decimal strings to dot-separated for parsing.
func (cf *CommaFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var rawValue string

	err := d.DecodeElement(&rawValue, &start)
	if err != nil {
		return fmt.Errorf("failed to decode XML element: %w", err)
	}

	rawValue = strings.ReplaceAll(rawValue, ",", ".")

	value, err := strconv.ParseFloat(rawValue, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float value: %w", err)
	}

	*cf = CommaFloat(value)

	return nil
}

// Settings holds configuration settings for input and output file paths.
type Settings struct {
	InputFileSetting  string `yaml:"input-file"`
	OutputFileSetting string `yaml:"output-file"`
}

// Valute represents a single currency exchange rate entry.
type Valute struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    CommaFloat `json:"value"     xml:"Value"`
}

// ValCurs represents a collection of currency exchange rates.
type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}
