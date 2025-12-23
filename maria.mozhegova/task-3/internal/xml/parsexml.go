package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html/charset"
)

func ParseXML(path string, result any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, encoding)
	}

	err = decoder.Decode(result)
	if err != nil {
		return fmt.Errorf("failed to parse XML: %w", err)
	}

	return nil
}
