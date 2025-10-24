package IOcurrency

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []struct {
		NumCode  int    `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		ValueStr string `xml:"Value"`
	} `xml:"Valute"`
}

type ValuteJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func (v *ValCurs) ReadXML(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)

	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.ToLower(charset) == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	if err := decoder.Decode(&v); err != nil {
		panic(err)
	}
}

func (v *ValCurs) Sort() {
	sort.Slice(v.Valutes, func(i, j int) bool {
		return v.Valutes[i].ValueStr > v.Valutes[j].ValueStr
	})
}

func SaveJSON(path string, data any) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "    ")
	return enc.Encode(data)
}
