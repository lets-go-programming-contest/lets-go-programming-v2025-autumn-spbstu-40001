package IOcurrency

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type ValCurs struct {
	XMLName xml.Name    `xml:"ValCurs"`
	Valutes []ValuteXML `xml:"Valute"`
}

type ValuteXML struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	ValueStr string `xml:"Value"`
}

type ValuteJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func LoadXML(path string) ([]ValuteJSON, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)

	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.ToLower(charset) == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	var curs ValCurs
	if err := decoder.Decode(&curs); err != nil {
		return nil, err
	}

	valutes := make([]ValuteJSON, 0, len(curs.Valutes))
	for _, v := range curs.Valutes {
		valueStr := strings.Replace(v.ValueStr, ",", ".", 1)

		var value float64
		if _, err := fmt.Sscanf(valueStr, "%f", &value); err != nil {
			return nil, fmt.Errorf("failed to parse value %s: %v", v.ValueStr, err)
		}

		valutes = append(valutes, ValuteJSON{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    value,
		})
	}
	return valutes, nil
}

func SaveJSON(path string, data []ValuteJSON) error {
	if err := os.MkdirAll(dir(path), 0755); err != nil {
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

func dir(path string) string {
	if i := len(path) - 1; i >= 0 {
		for i >= 0 && path[i] != '/' && path[i] != '\\' {
			i--
		}
		return path[:i]
	}
	return ""
}
