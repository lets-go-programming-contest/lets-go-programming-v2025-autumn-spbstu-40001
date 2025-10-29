package xmlparser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

func ParseXML(path string, result any) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read XML file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(file))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(result)
	if err != nil {
		return fmt.Errorf("failed to parse XML file: %w", err)
	}

	return nil
}
