package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	ID       string  `xml:"ID,attr"`
	NumCode  int     `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Nominal  int     `xml:"Nominal"`
	Name     string  `xml:"Name"`
	Value    float64 `xml:"Value"`
}

type currencyValue float64

func (cv *currencyValue) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var valueString string
	err := decoder.DecodeElement(&valueString, &start)
	if err != nil {
		return fmt.Errorf("decode element: %w", err)
	}

	valueString = strings.Replace(valueString, ",", ".", 1)

	value, err := strconv.ParseFloat(valueString, 64)
	if err != nil {
		return fmt.Errorf("parse float %q: %w", valueString, err)
	}

	*cv = currencyValue(value)

	return nil
}

func (v *Valute) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	type alias Valute

	var temp struct {
		alias
		Value currencyValue `xml:"Value"`
	}

	err := decoder.DecodeElement(&temp, &start)
	if err != nil {
		return fmt.Errorf("decode element: %w", err)
	}

	*v = Valute(temp.alias)
	v.Value = float64(temp.Value)

	return nil
}

func ParseXMLFile(inputFile string) (*ValCurs, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, encoding)
	}

	var valCurs ValCurs

	err = decoder.Decode(&valCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	return &valCurs, nil
}
