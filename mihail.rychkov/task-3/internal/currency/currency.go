package currency

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type Currency struct {
	XMLName  xml.Name `json:"-"         xml:"Valute"`
	NumCode  uint     `json:"num_code"  xml:"NumCode"`
	CharCode string   `json:"char_code" xml:"CharCode"`
	ValueStr string   `json:"-"         xml:"Value"`
	Value    float32  `json:"value"     xml:"-"`
}
type CurrencyRates struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Rates   []Currency `xml:"Valute"`
}

func (rates *CurrencyRates) Len() int {
	return len(rates.Rates)
}

func (rates *CurrencyRates) Less(i, j int) bool {
	return rates.Rates[i].Value < rates.Rates[j].Value
}

func (rates *CurrencyRates) Swap(i, j int) {
	rates.Rates[i], rates.Rates[j] = rates.Rates[j], rates.Rates[i]
}

func Prepare(rates *CurrencyRates) error {
	for idx := range len(rates.Rates) {
		value, err := strconv.ParseFloat(strings.ReplaceAll(rates.Rates[idx].ValueStr, ",", "."), 32)
		if err != nil {
			return fmt.Errorf("failed to parse rate value: %w", err)
		}

		rates.Rates[idx].Value = float32(value)
	}

	return nil
}

func ParseXml(xmlPath string) (CurrencyRates, error) {
	var result CurrencyRates

	xmlFile, err := os.Open(xmlPath)
	if err != nil {
		return result, fmt.Errorf("failed to open currency list xml file: %w", err)
	}

	defer func() {
		_ = xmlFile.Close()
	}()

	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&result)
	if err != nil {
		return result, fmt.Errorf("failed to parse currency list xml file: %w", err)
	}

	return result, nil
}
