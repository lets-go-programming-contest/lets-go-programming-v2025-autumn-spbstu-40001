package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

type FloatWithComma float64

func (f *FloatWithComma) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var valueStr string
	if err := d.DecodeElement(&valueStr, &start); err != nil {
		return err
	}

	cleaned := strings.Replace(valueStr, ",", ".", -1)
	floatVal, err := strconv.ParseFloat(cleaned, 64)
	if err != nil {
		return err
	}

	*f = FloatWithComma(floatVal)
	return nil
}

type ValCurs struct {
	Valutes []struct {
		NumCode  int            `xml:"NumCode" json:"num_code"`
		CharCode string         `xml:"CharCode" json:"char_code"`
		Value    FloatWithComma `xml:"Value" json:"value"`
	} `xml:"Valute"`
}

func parseXML(data []byte, v interface{}) error {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return input, nil
	}
	return decoder.Decode(v)
}

func main() {
	configPath := flag.String("config", "", "Path to yaml file")
	flag.Parse()

	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		panic(err)
	}

	configFile, err := os.ReadFile(*configPath)

	if err != nil {
		panic(err)
	}

	var config struct {
		InputFile  string `yaml:"input-file"`
		OutputFile string `yaml:"output-file"`
	}

	err = yaml.Unmarshal(configFile, &config) // &config - указатель на структуру
	if err != nil {
		panic(err)
	}

	inputFilePath := config.InputFile
	outputFilePath := config.OutputFile

	if _, err := os.Stat(inputFilePath); os.IsNotExist(err) {
		panic(err)
	}

	inputFile, err := os.ReadFile(inputFilePath)

	if err != nil {
		panic(err)
	}

	valCurs := new(ValCurs)
	err = parseXML(inputFile, valCurs)
	if err != nil {
		panic(err)
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	if err := os.MkdirAll(filepath.Dir(outputFilePath), 0755); err != nil {
		panic(err)
	}

	output_file, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}
	defer output_file.Close()

	encoder := json.NewEncoder(output_file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(valCurs.Valutes); err != nil {
		panic(err)
	}
}
