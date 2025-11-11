package iocurrency

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type UserFloat float64

func (f *UserFloat) UnmarshalText(text []byte) error {
	str := strings.ReplaceAll(string(text), ",", ".")
	val, err := strconv.ParseFloat(str, 64)

	if err != nil {

		return fmt.Errorf("parse float %q: %w", text, err)
	}

	*f = UserFloat(val)

	return nil
}

type ValCurs struct {
	Valutes []struct {
		NumCode  int       `json:"num_code"  xml:"NumCode"`
		CharCode string    `json:"char_code" xml:"CharCode"`
		ValueStr UserFloat `json:"value"     xml:"Value"`
	} `xml:"Valute"`
}

func (v *ValCurs) Sort() {
	sort.Slice(v.Valutes, func(i, j int) bool {
		return v.Valutes[i].ValueStr > v.Valutes[j].ValueStr
	})
}
