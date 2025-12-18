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
	Value    float64 `json:"value"     xml:"Value"`
}

type valCursXML struct {
	Valutes []struct {
		NumCode  int    `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	} `xml:"Valute"`
}

func ParseFromXMLFile(inputFilePath string) ([]Currency, error) {
	valCurs, err := xmlparser.ParseCurrencyRateFromXML[valCursXML](inputFilePath)
	if err != nil {
		return nil, err
	}

	currencies := make([]Currency, len(valCurs.Valutes))

	for i, v := range valCurs.Valutes {
		valueStr := strings.Replace(v.Value, ",", ".", 1)
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return nil, err
		}

		currencies[i] = Currency{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    value,
		}
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
