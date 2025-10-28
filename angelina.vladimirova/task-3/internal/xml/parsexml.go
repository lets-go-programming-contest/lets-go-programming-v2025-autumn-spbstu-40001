package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

func ParseXML(path string, result interface{}) error {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("cannot read xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(fileData))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(result)
	if err != nil {
		return fmt.Errorf("failed to parse xml file: %w", err)
	}

	return nil
}
