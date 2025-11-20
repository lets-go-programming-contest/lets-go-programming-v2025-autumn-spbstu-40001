package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int          `xml:"NumCode"`
	CharCode string       `xml:"CharCode"`
	Value    currencyValue `xml:"Value"`
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

func ConvertAndSort(valCurs *ValCurs) []Valute {
	currencies := make([]Valute, len(valCurs.Valutes))
	copy(currencies, valCurs.Valutes)

	sort.Slice(currencies, func(i, j int) bool {
		return float64(currencies[i].Value) > float64(currencies[j].Value)
	})

	return currencies
}
