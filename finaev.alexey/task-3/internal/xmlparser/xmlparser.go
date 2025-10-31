package xmlparser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html/charset"
)

func LoadCurrencies(inputFile string, res any) error {
	file, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(file))

	decoder.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, encoding)
	}

	if err := decoder.Decode(res); err != nil {
		return fmt.Errorf("error decode XML: %w", err)
	}

	return nil
}
