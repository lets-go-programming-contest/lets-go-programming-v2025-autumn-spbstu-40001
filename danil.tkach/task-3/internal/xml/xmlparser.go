package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	Valutes []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode    int     `json:"num_code"  xml:"NumCode"`
	CharCode   string  `json:"char_code" xml:"CharCode"`
	ValueField string  `json:"-"         xml:"Value"`
	Value      float64 `json:"value"     xml:"-"`
}

func ReadValCurs(inputFile string) (*ValCurs, error) {
	xmlData, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read XML-file %s: %w", inputFile, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))

	decoder.CharsetReader = func(c string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, c)
	}

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("failed to parse XML-file: %w", err)
	}

	for i := range valCurs.Valutes {
		valute := &valCurs.Valutes[i]
		valueStr := strings.Replace(valute.ValueField, ",", ".", 1)

		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("failed convert '%s' to number: %w", valute.ValueField, err)
		}

		valute.Value = value
	}

	return &valCurs, nil
}
