package xml

import (
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
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int     `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Value    float64 `xml:"-"`
}

type valuteRaw struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	ValueStr string `xml:"Value"`
}

func (v *Valute) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw valuteRaw
	if err := d.DecodeElement(&raw, &start); err != nil {
		return err
	}

	numCode, err := strconv.Atoi(raw.NumCode)
	if err != nil {
		return fmt.Errorf("parse NumCode '%s': %w", raw.NumCode, err)
	}

	valueStr := strings.Replace(raw.ValueStr, ",", ".", 1)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("parse Value '%s': %w", raw.ValueStr, err)
	}

	v.NumCode = numCode
	v.CharCode = raw.CharCode
	v.Value = value

	return nil
}

func ParseXML(filePath string) ([]Valute, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(label string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, label)
	}

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	return valCurs.Valutes, nil
}
