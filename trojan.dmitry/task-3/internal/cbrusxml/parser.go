package cbrusxml

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type FloatComma float64

func (f *FloatComma) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	if err := d.DecodeElement(&value, &start); err != nil {
		return err
	}

	value = strings.TrimSpace(value)
	value = strings.ReplaceAll(value, ",", ".")

	if value == "" {
		*f = 0
		return nil
	}

	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}

	*f = FloatComma(v)
	return nil
}

type Valute struct {
	NumCode  int        `xml:"NumCode"`
	CharCode string     `xml:"CharCode"`
	Nominal  int        `xml:"Nominal"`
	Value    FloatComma `xml:"Value"`
}

func ParseFile(path string) (*ValCurs, error) {
	val, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	dec := xml.NewDecoder(bytes.NewReader(val))

	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.ToLower(charset) == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	var vc ValCurs

	err = dec.Decode(&vc)
	if err != nil {
		return nil, err
	}

	return &vc, nil
}
