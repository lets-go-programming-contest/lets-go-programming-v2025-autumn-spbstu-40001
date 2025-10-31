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

type ParseFloat64 float64

func (c *ParseFloat64) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var valueStr string

	err := decoder.DecodeElement(&valueStr, &start)
	if err != nil {
		return fmt.Errorf("failed to parse value: %w", err)
	}

	// Заменяем запятую на точку и убираем пробелы
	valueStr = strings.TrimSpace(valueStr)
	valueStr = strings.Replace(valueStr, ",", ".", 1)

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("failed to parse value '%s': %w", valueStr, err)
	}

	*c = ParseFloat64(value)
	return nil
}

type Currency struct {
	NumCode  string       `xml:"NumCode"`
	CharCode string       `xml:"CharCode"`
	Value    ParseFloat64 `xml:"Value"`
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
