package xml

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html/charset"
)

type FileLoader interface {
	Load(filename string, target interface{}) error
}

type XMLLoader struct{}

func NewLoader() *XMLLoader {
	return &XMLLoader{}
}

func (l *XMLLoader) Load(filename string, target interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(label string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, label)
	}

	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("decode XML: %w", err)
	}

	return nil
}
