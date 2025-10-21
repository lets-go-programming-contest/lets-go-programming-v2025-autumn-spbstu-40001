package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
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
	Valute []struct {
		NumCode  string `xml:"NumCode" json:"num_code"`
		CharCode string `xml:"CharCode" json:"char_code"`
		Value    string `xml:"Value" json:"value"`
	}
}

var configPathFlag = flag.String("config", "config.yaml", "path to config file")

type ByValue ValCurs

func (myValCurs ByValue) Len() int {
	return len(myValCurs.Valute)
}

func (myValCurs ByValue) Swap(i, j int) {
	myValCurs.Valute[i], myValCurs.Valute[j] = myValCurs.Valute[j], myValCurs.Valute[i]
}

func (myValCurs ByValue) Less(i, j int) bool {
	return myValCurs.Valute[i].Value > myValCurs.Valute[j].Value
}

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
		errCreate := os.Mkdir(outputFileInfo.Dir, 0777)
		if errCreate != nil {
			panic("cannot make directory")
		}
		_, errCreate = os.OpenFile(path.Join(outputFileInfo.Dir, outputFileInfo.Filename), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
		if errCreate != nil {
			panic("cannot create file")
		}
	}

	inputFile, err := os.Open(files.InputFile)
	if err != nil && os.IsNotExist(err) {
		panic("input file does not exist")
	}

	XMLDecoder := xml.NewDecoder(inputFile)

	var CentroBankValuteCourses ValCurs

	if err := XMLDecoder.Decode(&CentroBankValuteCourses); err != nil {
		panic(err)
	}

	for _, v := range CentroBankValuteCourses.Valute {
		fmt.Println(v.NumCode)
		fmt.Println(v.CharCode)
		fmt.Println(v.Value)
	}

	sort.Sort(ByValue(CentroBankValuteCourses))

	outputFile, err := os.OpenFile(path.Join(outputFileInfo.Dir, outputFileInfo.Filename), os.O_WRONLY, 0777)

	if err != nil {
		panic(err)
	}

	JSONEncoder := json.NewEncoder(outputFile)
	JSONEncoder.SetIndent("", " ")

	JSONEncoder.Encode(&CentroBankValuteCourses.Valute)
}
