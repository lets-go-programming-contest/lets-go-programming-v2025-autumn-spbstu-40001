package currency

import (
	"fmt"
	"strconv"
	"strings"
)

type Converter interface {
	Convert(source interface{}) (interface{}, error)
}

type CurrencyConverter struct{}

func NewConverter() *CurrencyConverter {
	return &CurrencyConverter{}
}

type XMLValute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type XMLValCurs struct {
	Valutes []XMLValute `xml:"Valute"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type CurrencyCollection []Currency

func (c CurrencyCollection) Len() int           { return len(c) }
func (c CurrencyCollection) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c CurrencyCollection) Less(i, j int) bool { return c[i].Value > c[j].Value }

func (conv *CurrencyConverter) Convert(source interface{}) (interface{}, error) {
	valCurs, ok := source.(*XMLValCurs)
	if !ok {
		return nil, fmt.Errorf("invalid source type for currency conversion")
	}

	currencies := make(CurrencyCollection, 0, len(valCurs.Valutes))

	for _, valute := range valCurs.Valutes {
		currency, err := conv.convertValute(valute)
		if err != nil {
			return nil, err
		}
		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (conv *CurrencyConverter) convertValute(valute XMLValute) (Currency, error) {
	numCode, err := strconv.Atoi(valute.NumCode)
	if err != nil {
		return Currency{}, fmt.Errorf("parse NumCode '%s': %w", valute.NumCode, err)
	}

	valueStr := strings.Replace(valute.Value, ",", ".", 1)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return Currency{}, fmt.Errorf("parse Value '%s': %w", valute.Value, err)
	}

	return Currency{
		NumCode:  numCode,
		CharCode: valute.CharCode,
		Value:    value,
	}, nil
}
