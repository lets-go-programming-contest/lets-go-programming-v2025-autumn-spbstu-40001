package xmlparser

import (
	"bytes"
	"encoding/xml"
	"os"

	"golang.org/x/net/html/charset"
)

func Parse(path string, result any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(result)
	if err != nil {
		return err
	}

	return nil
}
