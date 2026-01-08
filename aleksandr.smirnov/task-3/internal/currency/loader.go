package currency

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	if charset == "windows-1251" {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	return input, nil
}

func DecodeXML(r io.Reader) (*Bank, error) {
	decoder := xml.NewDecoder(r)
	decoder.CharsetReader = charsetReader

	var data Bank
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("xml decode failed: %w", err)
	}

	return &data, nil
}

func LoadFromFile(path string) (*Bank, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}

	bankData, decodeErr := DecodeXML(file)

	if closeErr := file.Close(); closeErr != nil && decodeErr == nil {
		return nil, fmt.Errorf("failed to close file: %w", closeErr)
	}

	return bankData, decodeErr
}
