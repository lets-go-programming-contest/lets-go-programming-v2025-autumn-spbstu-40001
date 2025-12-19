package currency

import (
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/vikaglushkova/task-3/internal/xmlparser"
)

type xmlFloat float64

func (xf *xmlFloat) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var content string

	err := decoder.DecodeElement(&content, &start)
	if err != nil {
		return fmt.Errorf("cannot decode element: %w", err)
	}

	content = strings.Replace(content, ",", ".", 1)

	value, err := strconv.ParseFloat(content, 64)
	if err != nil {
		return fmt.Errorf("cannot parse float %s: %w", content, err)
	}

	*xf = xmlFloat(value)

	return nil
}

type Currency struct {
	NumCode  int      `json:"num_code"  xml:"NumCode"`
	CharCode string   `json:"char_code" xml:"CharCode"`
	Value    xmlFloat `json:"value"     xml:"Value"`
}

type ValCursXML struct {
	Valutes []Currency `xml:"Valute"`
}

func ParseFromXMLFile(inputFilePath string) ([]Currency, error) {
	valCurs, err := xmlparser.ParseCurrencyRateFromXML[ValCursXML](inputFilePath)
	if err != nil {
		return nil, err
	}

	return valCurs.Valutes, nil
}

func ConvertAndSort(currencies []Currency) []Currency {
	result := make([]Currency, len(currencies))
	copy(result, currencies)

	sort.Slice(result, func(i, j int) bool {
		return float64(result[i].Value) > float64(result[j].Value)
	})

	return result
}
