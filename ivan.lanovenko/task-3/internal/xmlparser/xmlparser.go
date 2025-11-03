package xmlparser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"

	"golang.org/x/net/html/charset"
)

func ParseXML[T any](data []byte) (*T, error) {
	var result T

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(charSet string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, charSet)
	}

	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	return &result, nil
}
