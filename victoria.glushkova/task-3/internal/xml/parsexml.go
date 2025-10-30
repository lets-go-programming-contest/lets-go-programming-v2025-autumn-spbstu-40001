package xml

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	ID       string  `xml:"ID,attr"`
	NumCode  int     `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Nominal  int     `xml:"Nominal"`
	Name     string  `xml:"Name"`
	Value    float64 `xml:"Value"`
}

func ParseXMLFile(inputFile string) (*ValCurs, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open input file: %w", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		_ = file.Close()
		return nil, fmt.Errorf("cannot read input file: %w", err)
	}

	_ = file.Close()

	var valCurs ValCurs
	if err = xml.Unmarshal(data, &valCurs); err != nil {
		return nil, fmt.Errorf("cannot parse XML data: %w", err)
	}

	return &valCurs, nil
}
