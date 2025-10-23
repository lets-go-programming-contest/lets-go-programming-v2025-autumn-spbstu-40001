package IOcurrency

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
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

	var curs ValCurs
	if err := xml.NewDecoder(file).Decode(&curs); err != nil {
		return nil, err
	}

	valutes := make([]ValuteJSON, 0, len(curs.Valutes))
	for _, v := range curs.Valutes {
		var value float64
		_, err := fmt.Sscanf(v.ValueStr, "%f", &value)
		if err != nil {
			fmt.Sscanf(replaceComma(v.ValueStr), "%f", &value)
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

func replaceComma(s string) string {
	out := []rune{}
	for _, r := range s {
		if r == ',' {
			out = append(out, '.')
		} else {
			out = append(out, r)
		}
	}
	return string(out)
}
