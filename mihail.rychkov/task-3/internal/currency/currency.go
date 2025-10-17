package currency

import (
	"bytes"
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
	NumCode  uint    `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    float32 `json:"value"     xml:"-"`
}
type CurrencyRates struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Rates   []Currency `xml:"Valute"`
}

func (currency *Currency) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	type Datatype Currency

	type Wrapper struct {
		*Datatype

		ValueStr string `xml:"Value"`
	}

	wrapper := Wrapper{Datatype: (*Datatype)(currency), ValueStr: ""}

	err := decoder.DecodeElement(&wrapper, &start)
	if err != nil {
		return fmt.Errorf("failed to parse raw Valute: %w", err)
	}

	value, err := strconv.ParseFloat(strings.Replace(wrapper.ValueStr, ",", ".", 1), 32)
	if err != nil {
		return fmt.Errorf("failed to parse rate value: %w", err)
	}

	currency.Value = float32(value)

	return nil
}

func ParseXML(xmlPath string) (CurrencyRates, error) {
	var result CurrencyRates

	xmlFileData, err := os.ReadFile(xmlPath)
	if err != nil {
		return result, fmt.Errorf("cannot read currency list xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(xmlFileData))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&result)
	if err != nil {
		return result, fmt.Errorf("failed to parse currency list xml file: %w", err)
	}

	return result, nil
}

func ForceWriteToJSON(rates *CurrencyRates, outPath string, defaultMode os.FileMode) error {
	serialized, err := json.MarshalIndent(rates.Rates, "", "\t")
	if err != nil {
		return fmt.Errorf("failed to serialize data to json: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(outPath), defaultMode)
	if err != nil {
		return fmt.Errorf("cannot create required directories: %w", err)
	}

	err = os.WriteFile(outPath, serialized, defaultMode)
	if err != nil {
		return fmt.Errorf("cannot write output file: %w", err)
	}

	return nil
}
