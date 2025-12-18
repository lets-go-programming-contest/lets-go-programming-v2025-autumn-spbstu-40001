package currency

import (
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/vikaglushkova/task-3/internal/xmlparser"
)

type Currency struct {
	NumCode  int     `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    float64 `json:"value"     xml:"Value"`
}

type xmlCurrency Currency

func (xc *xmlCurrency) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	type rawCurrency struct {
		NumCode  int    `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	}

	var raw rawCurrency

	if err := decoder.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode element: %w", err)
	}

	valueStr := strings.Replace(raw.Value, ",", ".", 1)
	value, err := strconv.ParseFloat(valueStr, 64)

	if err != nil {
		return fmt.Errorf("parse value %q: %w", raw.Value, err)
	}

	xc.NumCode = raw.NumCode
	xc.CharCode = raw.CharCode
	xc.Value = value

	return nil
}

type ValCursXML struct {
	Valutes []xmlCurrency `xml:"Valute"`
}

func ParseFromXMLFile(inputFilePath string) ([]Currency, error) {
	valCurs, err := xmlparser.ParseCurrencyRateFromXML[ValCursXML](inputFilePath)
	if err != nil {
		return nil, err
	}

	currencies := make([]Currency, len(valCurs.Valutes))

	for i, xc := range valCurs.Valutes {
		currencies[i] = Currency(xc)
	}

	return currencies, nil
}

func ConvertAndSort(currencies []Currency) []Currency {
	result := make([]Currency, len(currencies))
	copy(result, currencies)

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}
