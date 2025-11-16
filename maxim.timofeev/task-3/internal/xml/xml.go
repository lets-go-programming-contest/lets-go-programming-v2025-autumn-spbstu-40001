package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html/charset"
)

func NewDecoder(r io.Reader) *xml.Decoder {
	decoder := xml.NewDecoder(r)
	decoder.CharsetReader = charset.NewReaderLabel

	return decoder
}

func ReadXML(path string, out any) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read XML file: %w", err)
	}

	decoder := NewDecoder(bytes.NewReader(file))

	if err := decoder.Decode(out); err != nil {
		return fmt.Errorf("decode XML: %w", err)
	}

	return nil
}
