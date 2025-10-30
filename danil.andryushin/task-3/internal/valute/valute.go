package valute

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type (
	ValuteValue float32
	Valute      struct {
		NumCode  int         `xml:"NumCode" json:"num_code"`
		CharCode string      `xml:"CharCode" json:"char_code"`
		Value    ValuteValue `xml:"Value" json:"value"`
	}
	ValuteCourse struct { // TODO: renaming
		XMLName xml.Name `xml:"ValCurs"`
		Valutes []Valute `xml:"Valute"`
	}
)

func (obj *ValuteValue) UnmarshalXML(decode *xml.Decoder, start xml.StartElement) error {
	var value string
	err := decode.DecodeElement(&value, &start)
	if err != nil {
		return err
	}
	value = strings.ReplaceAll(value, ",", ".")
	temp, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return err
	}
	*obj = ValuteValue(temp)
	return nil
}
