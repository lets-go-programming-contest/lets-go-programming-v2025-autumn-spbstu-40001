package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

const (
	DirPerm  = 0o755
	FilePerm = 0o644
)

type ValCurs struct {
	Valutes []struct {
		NumCode  int           `json:"num_code"  xml:"NumCode"`
		CharCode string        `json:"char_code" xml:"CharCode"`
		Value    CustomFloat64 `json:"value"     xml:"Value"`
	} `xml:"Valute"`
}

type CustomFloat64 float64

func (c *CustomFloat64) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var valueStr string

	err := decoder.DecodeElement(&valueStr, &start)
	if err != nil {
		return fmt.Errorf("failed to parse value: %w", err)
	}

	value, err := strconv.ParseFloat(strings.Replace(valueStr, ",", ".", 1), 64)
	if err != nil {
		return fmt.Errorf("failed to parse value: %w", err)
	}

	*c = CustomFloat64(value)

	return nil
}

func ReadValCurs(path string) (*ValCurs, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, encoding)
	}

	var valCurs ValCurs

	err = decoder.Decode(&valCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	return &valCurs, nil
}

func (v *ValCurs) SortByValueDesc() {
	sort.Slice(v.Valutes, func(i, j int) bool {
		return v.Valutes[i].Value > v.Valutes[j].Value
	})
}

func WriteToJSON(valCurs *ValCurs, path string) error {
	err := os.MkdirAll(filepath.Dir(path), DirPerm)
	if err != nil {
		return fmt.Errorf("failed to create a dir: %w", err)
	}

	data, err := json.MarshalIndent(valCurs.Valutes, "", " ")
	if err != nil {
		return fmt.Errorf("failed to write JSON: %w", err)
	}

	err = os.WriteFile(path, data, FilePerm)
	if err != nil {
		return fmt.Errorf("failed to write JSON: %w", err)
	}

	return nil
}
