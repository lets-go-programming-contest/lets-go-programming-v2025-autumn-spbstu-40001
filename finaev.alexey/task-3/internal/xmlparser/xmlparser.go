package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html/charset"
)

func LoadCurrencies(inputFile string, res any) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed open file: %w", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)

	decoder.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, encoding)
	}

	if err := decoder.Decode(res); err != nil {
		return fmt.Errorf("error decode XML: %w", err)
	}

	return nil
}
