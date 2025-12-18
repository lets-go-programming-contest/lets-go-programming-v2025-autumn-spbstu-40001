package xmlparser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html/charset"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseCurrencyRateFromXML[T any](inputFilePath string) (*T, error) {
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("can't read file %s: %w", inputFilePath, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, encoding)
	}

	var result T
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode XML: %w", err)
	}

	return &result, nil
}
