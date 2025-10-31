package iocurrency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type CurrencyList struct {
	XMLName xml.Name `xml:"ValCurs"`
	Items   []struct {
		NumericCode int     `json:"num_code"  xml:"NumCode"`
		Code        string  `json:"char_code" xml:"CharCode"`
		Rate        float64 `json:"value"     xml:"Value"`
	} `xml:"Valute"`
}

func (cl *CurrencyList) ParseXML(filepath string) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(bytes.ReplaceAll(data, []byte(","), []byte(".")))
	parser := xml.NewDecoder(reader)

	parser.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		if strings.ToLower(encoding) == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return input, nil
	}

	if err := parser.Decode(&cl); err != nil {
		panic(err)
	}
}

func (cl *CurrencyList) OrderByValue() {
	sort.Slice(cl.Items, func(i, j int) bool {
		return cl.Items[i].Rate > cl.Items[j].Rate
	})
}

func ExportJSON(path string, data interface{}) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
