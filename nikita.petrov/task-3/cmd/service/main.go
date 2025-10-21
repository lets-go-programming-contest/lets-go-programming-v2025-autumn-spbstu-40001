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

	var dir string
	var filename string

	if strings.Contains(files.OutputFile, "/") {
		outputFilePath := strings.Split(files.OutputFile, "/")
		dir = outputFilePath[0]
		filename = outputFilePath[1]
	} else {
		filename = files.OutputFile
	}

	_, err = os.Stat(files.OutputFile)
	if errors.Is(err, os.ErrNotExist) {
		if dir != "" {
			errCreateDir := os.Mkdir(dir, 0777)
			if errCreateDir != nil {
				panic("cannot make directory")
			}
		}
		_, errCreateFile := os.OpenFile(path.Join(dir, filename), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
		if errCreateFile != nil {
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

	outputFile, err := os.OpenFile(path.Join(dir, filename), os.O_WRONLY, 0777)

	if err != nil {
		panic(err)
	}

	JSONEncoder := json.NewEncoder(outputFile)
	JSONEncoder.SetIndent("", "\t")

	JSONEncoder.Encode(&CentroBankValuteCourses.Valute)
}
