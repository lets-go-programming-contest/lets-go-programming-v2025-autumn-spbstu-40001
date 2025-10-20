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
		NumCode  int     `xml:"NumCode" json:"num_code"`
		CharCode string  `xml:"CharCode" json:"char_code"`
		Value    float64 `xml:"Value" json:"value"`
	} `xml:"Valute"`
}

func (v *ValCurs) ParseXML(data []byte) {
	decoder := xml.NewDecoder(bytes.NewReader(bytes.Replace(data, []byte(","), []byte("."), -1)))
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
