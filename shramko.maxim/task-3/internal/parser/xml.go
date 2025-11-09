package parser

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"strings"

	"github.com/Elektrek/task-3/internal/model"

	"golang.org/x/net/html/charset"
)

func ParseCurrencies(filepath string) (*model.CurrencyCollection, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(bytes.ReplaceAll(data, []byte(","), []byte(".")))
	decoder := xml.NewDecoder(reader)

	decoder.CharsetReader = func(charsetLabel string, input io.Reader) (io.Reader, error) {
		encoding := strings.ToLower(charsetLabel)
		if encoding == "" || encoding == "utf-8" || encoding == "utf8" {
			return input, nil
		}

		if encodingHandler, found := charset.Lookup(encoding); found {
			return encodingHandler.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	var result struct {
		XMLName    xml.Name         `xml:"ValCurs"`
		Currencies []model.Currency `xml:"Valute"`
	}

	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}

	return &model.CurrencyCollection{Currencies: result.Currencies}, nil
}
