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
)

type ValCurs struct {
	Valutes []struct {
		NumCode  string `xml:"NumCode" json:"num_code"`
		CharCode string `xml:"CharCode" json:"char_code"`
		Value    string `xml:"Value" json:"value"`
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

	lines := strings.Split(string(configFile), "\n")

	inputFilePath := lines[0][strings.Index((lines[0]), "\"")+1 : strings.LastIndex((lines[0]), "\"")]
	outputFilePath := lines[1][strings.Index((lines[1]), "\"")+1 : strings.LastIndex((lines[1]), "\"")]

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

	for i := range valCurs.Valutes {
		valCurs.Valutes[i].Value = strings.Replace(valCurs.Valutes[i].Value, ",", ".", -1)
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		valI, _ := strconv.ParseFloat(valCurs.Valutes[i].Value, 64)
		valJ, _ := strconv.ParseFloat(valCurs.Valutes[j].Value, 64)
		return valI > valJ
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
