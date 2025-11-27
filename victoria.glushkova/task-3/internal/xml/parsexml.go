package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/vikaglushkova/task-3/internal/currency"
	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	Valutes []currency.Currency `xml:"Valute"`
}

func ParseXMLFile(inputFile string) (*ValCurs, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, encoding)
	}

	var valCurs ValCurs

	err = decoder.Decode(&valCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	return &valCurs, nil
}
