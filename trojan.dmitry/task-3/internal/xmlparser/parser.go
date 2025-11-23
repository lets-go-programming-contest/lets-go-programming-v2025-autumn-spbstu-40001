package xmlparser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html/charset"
)

func ParseXMLFile[T any](path string) (*T, error) {
	var dest T
	if err := ParseXMLFileInto(path, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

func ParseXMLFileInto(path string, v any) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file %s: %w", path, err)
	}

	dec := xml.NewDecoder(bytes.NewReader(raw))

	dec.CharsetReader = charset.NewReaderLabel

	if err := dec.Decode(v); err != nil && err != io.EOF {
		return fmt.Errorf("decode xml: %w", err)
	}
	return nil
}
