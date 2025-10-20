package valcurs

import (
	"bytes"
	"encoding/xml"
	"io"
	"sort"

	"golang.org/x/text/encoding/charmap"
)

type ValCurs struct {
	Valutes []struct {
		NumCode  int     `json:"num_code"  xml:"NumCode"`
		CharCode string  `json:"char_code" xml:"CharCode"`
		Value    float64 `json:"value"     xml:"Value"`
	} `xml:"Valute"`
}

func (v *ValCurs) ParseXML(data []byte) {
	decoder := xml.NewDecoder(bytes.NewReader(bytes.ReplaceAll(data, []byte(","), []byte("."))))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	if err := decoder.Decode(v); err != nil {
		panic(err)
	}
}

func (v *ValCurs) SortByValueDown() {
	sort.Slice(v.Valutes, func(i, j int) bool {
		return v.Valutes[i].Value > v.Valutes[j].Value
	})
}
