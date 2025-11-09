package xmlparser

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"strings"

	"github.com/Elektrek/task-3/internal/model"

	"golang.org/x/net/html/charset"
)

func ParseXML(filepath string) (*model.CurrencyList, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(bytes.ReplaceAll(data, []byte(","), []byte(".")))
	parser := xml.NewDecoder(reader)

	parser.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		enc := strings.ToLower(encoding)
		if enc == "" || enc == "utf-8" || enc == "utf8" {
			return input, nil
		}

		if alias, ok := charset.Lookup(enc); ok {
			return alias.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	var currencyData struct {
		XMLName xml.Name             `xml:"ValCurs"`
		Items   []model.CurrencyItem `xml:"Valute"`
	}

	if err := parser.Decode(&currencyData); err != nil {
		return nil, err
	}

	return &model.CurrencyList{Items: currencyData.Items}, nil
}
