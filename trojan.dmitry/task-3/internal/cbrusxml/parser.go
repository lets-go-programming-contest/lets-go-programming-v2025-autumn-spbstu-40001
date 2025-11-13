package cbrusxml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type FloatComma float64

func (f *FloatComma) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	if err := d.DecodeElement(&value, &start); err != nil {
		return fmt.Errorf("decode element: %w", err)
	}

	value = strings.TrimSpace(value)
	value = strings.ReplaceAll(value, ",", ".")

	if value == "" {
		*f = 0

		return nil
	}

	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("parse float: %w", err)
	}

	*f = FloatComma(val)

	return nil
}

type Valute struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    FloatComma `json:"value"     xml:"Value"`
}

func ParseFile(path string) (*ValCurs, error) {
	val, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	dec := xml.NewDecoder(bytes.NewReader(val))

	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch strings.ToLower(charset) {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		case "utf-8":
			return input, nil
		default:
			return input, nil
		}
	}

	var valC ValCurs

	err = dec.Decode(&valC)
	if err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	return &valC, nil
}
