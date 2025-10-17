package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

func ParseFile[T any](path string, result *T) error {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("cannot read currency list xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(fileData))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(result)
	if err != nil {
		return fmt.Errorf("failed to parse currency list xml file: %w", err)
	}

	return nil
}
