package currency

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
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

func ParseXML(xmlPath string) (CurrencyRates, error) {
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

func ForceWriteAsJson(rates *CurrencyRates, outPath string, defaultMode os.FileMode) error {
	serialized, err := json.MarshalIndent(rates.Rates, "", "\t")
	if err != nil {
		return fmt.Errorf("failed to serialize data to json: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(outPath), os.ModeDir|defaultMode)
	if err != nil {
		return fmt.Errorf("failed to make required directories: %w", err)
	}

	err = os.WriteFile(outPath, append(serialized, '\n'), defaultMode)
	if err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}
