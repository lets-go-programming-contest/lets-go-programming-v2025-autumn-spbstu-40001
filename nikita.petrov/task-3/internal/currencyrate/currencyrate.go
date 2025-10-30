package currencyrate

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"
)

var ( 
	ErrElemDecode error = errors.New("cannot decode element")
    ErrFloatParse error = errors.New("cannot parse float")
)

type floatWithDots float64

type CurrencyRate struct {
	Valutes []*singleValute `json:"valute" xml:"Valute"`
}

type singleValute struct {
	NumCode  int           `json:"num_code"  xml:"NumCode"`
	CharCode string        `json:"char_code" xml:"CharCode"`
	Value    floatWithDots `json:"value"     xml:"Value"`
}

func (fd *floatWithDots) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string

	err := d.DecodeElement(&content, &start)
	if err != nil {
		return ErrElemDecode
	}

	content = strings.ReplaceAll(content, ",", ".")

	var retFloat64 float64

	retFloat64, err = strconv.ParseFloat(content, 64)
	if err != nil {
		return ErrFloatParse
	}

	*fd = floatWithDots(retFloat64)

	return nil
}
