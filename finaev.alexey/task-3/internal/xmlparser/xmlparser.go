package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"

	"golang.org/x/net/html/charset"
)

type Currency struct {
	NumCode  string  `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Value    float64 `xml:"Value"`
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
