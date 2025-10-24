package IOcurrency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
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
		NumCode  int     `json:"num_code" xml:"NumCode"`
		CharCode string  `json:"char_code" xml:"CharCode"`
		ValueStr float64 `json:"value" xml:"Value"`
	} `xml:"Valute"`
}

func (v *ValCurs) ReadXML(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(bytes.ReplaceAll(file, []byte(","), []byte("."))))

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
		return fmt.Errorf("failed to create directory for %s: %w", path, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", path, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "    ")

	return enc.Encode(data)
}
