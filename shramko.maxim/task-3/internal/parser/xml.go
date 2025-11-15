package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Elektrek/task-3/internal/model"
	"golang.org/x/text/encoding/charmap"
)

func ParseCurrencies(filePath string) (*model.CurrencyCollection, error) {
	fileContent, readErr := os.ReadFile(filePath)
	if readErr != nil {
		return nil, fmt.Errorf("failed to read file: %w", readErr)
	}

	dataReader := bytes.NewReader(bytes.ReplaceAll(fileContent, []byte(","), []byte(".")))
	xmlDecoder := xml.NewDecoder(dataReader)

	xmlDecoder.CharsetReader = func(encodingName string, stream io.Reader) (io.Reader, error) {
		enc := strings.ToLower(encodingName)

		switch enc {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(stream), nil
		case "utf-8", "utf8", "":
			return stream, nil
		default:
			return stream, nil
		}
	}

	var parsedData struct {
		XMLName    xml.Name         `xml:"ValCurs"`
		Currencies []model.Currency `xml:"Valute"`
	}

	if decodeErr := xmlDecoder.Decode(&parsedData); decodeErr != nil {
		return nil, fmt.Errorf("failed to decode XML: %w", decodeErr)
	}

	return &model.CurrencyCollection{CurrencyItems: parsedData.Currencies}, nil
}
