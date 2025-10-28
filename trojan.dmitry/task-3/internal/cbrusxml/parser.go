package cbrusxml

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml: "Date, attr"`
	Name    string   `xml: "name, attr"`
	Valutes []Valute `xml:"valute"`
}

type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

func ParseFile(path string) (*ValCurs, error) {
	val, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	normalized := bytes.ReplaceAll(val, []byte(","), []byte("."))
	dec := xml.NewDecoder(bytes.NewReader(normalized))

	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.EqualFold(strings.TrimSpace(charset), "windows-1251") {
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
