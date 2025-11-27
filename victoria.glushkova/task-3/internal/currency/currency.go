package currency

import (
        "encoding/xml"
        "fmt"
        "sort"
        "strconv"
        "strings"
)

type Currency struct {
        NumCode  int          `json:"num_code"  xml:"NumCode"`
        CharCode string       `json:"char_code" xml:"CharCode"`
        Value    CurrencyValue `json:"value"     xml:"Value"`
}

type CurrencyValue float64

func (cv *CurrencyValue) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
        var valueString string

        err := decoder.DecodeElement(&valueString, &start)
        if err != nil {
                return fmt.Errorf("decode element: %w", err)
        }

        valueString = strings.Replace(valueString, ",", ".", 1)

        value, err := strconv.ParseFloat(valueString, 64)
        if err != nil {
                return fmt.Errorf("parse float %q: %w", valueString, err)
        }

        *cv = CurrencyValue(value)

        return nil
}

func ConvertAndSort(currencies []Currency) []Currency {
        result := make([]Currency, len(currencies))
        copy(result, currencies)

        sort.Slice(result, func(i, j int) bool {
                return float64(result[i].Value) > float64(result[j].Value)
        })

        return result
}
