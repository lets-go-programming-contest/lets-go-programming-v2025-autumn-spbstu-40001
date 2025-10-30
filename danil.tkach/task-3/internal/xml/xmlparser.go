package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html/charset"
)

func DecodeXMLFile[T any](inputFile string, target *T) error {
	xmlData, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read XML-file %s: %w", inputFile, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))

	decoder.CharsetReader = func(c string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, c)
	}

	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("failed to parse XML-file: %w", err)
	}

	return nil
}
