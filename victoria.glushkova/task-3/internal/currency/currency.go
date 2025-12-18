package currency

import (
	"sort"
	"strconv"
	"strings"

	"github.com/vikaglushkova/task-3/internal/xmlparser"
)

type Currency struct {
	NumCode  int     `json:"num_code" xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    float64 `json:"value" xml:"Value"`
}

type ValueWithComma float64

func (v *ValueWithComma) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var valueStr string
	if err := d.DecodeElement(&valueStr, &start); err != nil {
		return err
	}

	valueStr = strings.Replace(valueStr, ",", ".", 1)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return err
	}

	*v = ValueWithComma(value)
	return nil
}

type currencyXML struct {
	NumCode  int           `xml:"NumCode"`
	CharCode string        `xml:"CharCode"`
	Value    ValueWithComma `xml:"Value"`
}

func (c currencyXML) ToCurrency() Currency {
	return Currency{
		NumCode:  c.NumCode,
		CharCode: c.CharCode,
		Value:    float64(c.Value),
	}
}

type ValCursXML struct {
	Valutes []currencyXML `xml:"Valute"`
}

func ParseFromXMLFile(inputFilePath string) ([]Currency, error) {
	valCurs, err := xmlparser.ParseCurrencyRateFromXML[ValCursXML](inputFilePath)
	if err != nil {
		return nil, err
	}

	currencies := make([]Currency, len(valCurs.Valutes))
	for i, xmlCurr := range valCurs.Valutes {
		currencies[i] = xmlCurr.ToCurrency()
	}

	return currencies, nil
}

func ConvertAndSort(currencies []Currency) []Currency {
	result := make([]Currency, len(currencies))
	copy(result, currencies)

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}
