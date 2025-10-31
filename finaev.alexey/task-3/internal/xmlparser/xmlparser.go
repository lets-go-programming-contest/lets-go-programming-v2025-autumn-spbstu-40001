package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type Currency struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    float64
}

func (c *Currency) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type currencyXML struct {
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	}

	var temp currencyXML
	if err := d.DecodeElement(&temp, &start); err != nil {
		return err
	}

	c.NumCode = temp.NumCode
	c.CharCode = temp.CharCode

	valueStr := strings.Replace(temp.Value, ",", ".", -1)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("Error parsing value '%s': %w", temp.Value, err)
	}
	c.Value = value

	return nil
}

type ValCurs struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}

func LoadCurrencies(inputFile string) (*ValCurs, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("Failed open file: %w", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)

	decoder.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, encoding)
	}

	var valCurs ValCurs

	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("Error decode XML: %w", err)
	}

	return &valCurs, nil
}

func (v *ValCurs) SortCurrenciesByValueDesc() {
	sort.Slice(v.Currencies, func(i, j int) bool {
		return v.Currencies[i].Value > v.Currencies[j].Value
	})
}

func (v *ValCurs) GetCurrencies() []Currency {
	return v.Currencies
}
