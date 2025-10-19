package main

import (
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type OutputFileInfo struct {
	Dir      string
	Filename string
}

func New(outputFile string) OutputFileInfo {
	outputFilePath := strings.Split(outputFile, "/")
	return OutputFileInfo{outputFilePath[0], outputFilePath[1]}
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Text    xml.CharData
	Valute  []struct {
		Text      xml.CharData
		NumCode   string `xml:"NumCode"`
		CharCode  string `xml:"CharCode"`
		Nominal   string `xml:"Nominal"`
		Name      string `xml:"Name"`
		Value     string `xml:"Value"`
		VunitRate string `xml:"VunitRate"`
	}
}

var configPathFlag = flag.String("config", "config.yaml", "path to config file")

func main() {
	configFile, err := os.Open(*configPathFlag)
	if err != nil && os.IsNotExist(err) {
		panic("config file does not exist")
	}

	configData, err := io.ReadAll(configFile)
	if err != nil {
		panic("cannot read file")
	}

	var files Config

	err = yaml.Unmarshal([]byte(configData), &files)

	if err != nil {
		panic("cannot unmarshal config data")
	}

	outputFileInfo := New(files.OutputFile)

	_, err = os.Stat(files.OutputFile)
	if errors.Is(err, os.ErrNotExist) {
		errMkDir := os.Mkdir(outputFileInfo.Dir, 0777)
		if errMkDir != nil {
			panic("cannot make directory")
		}
		_, errMkFile := os.OpenFile(path.Join(outputFileInfo.Dir, outputFileInfo.Filename), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
		if errMkFile != nil {
			panic("cannot create file")
		}
	}

	inputFile, err := os.Open(files.InputFile)
	if err != nil && os.IsNotExist(err) {
		panic("input file does not exist")
	}

	decoder := xml.NewDecoder(inputFile)

	var CentroBankValuteCourses ValCurs

	if err := decoder.Decode(&CentroBankValuteCourses); err != nil {
		panic(err)
	}
	fmt.Println(CentroBankValuteCourses)
}
