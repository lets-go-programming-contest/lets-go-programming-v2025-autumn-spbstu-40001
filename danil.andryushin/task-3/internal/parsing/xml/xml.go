package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

func Unmarshall(path string, obj any) error {
	input, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(input))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&obj)
	if err != nil {
		return fmt.Errorf("unmarshalling failed: %w", err)
	}

	return nil
}
