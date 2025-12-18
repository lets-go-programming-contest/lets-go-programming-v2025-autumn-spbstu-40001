package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"golang.org/x/net/html/charset"
)

func ParseXMLFile[T any](inputFile string) (*T, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, encoding)
	}

	var result T
	err = decoder.Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	return &result, nil
}
