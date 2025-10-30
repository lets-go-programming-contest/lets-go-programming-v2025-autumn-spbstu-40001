package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"os"
	"path"
	"sort"

	"github.com/Nekich06/task-3/internal/charsetsetter"
	"github.com/Nekich06/task-3/internal/config"
	"github.com/Nekich06/task-3/internal/currencyrate"
	"github.com/Nekich06/task-3/internal/filesmanager"
	"github.com/Nekich06/task-3/internal/valutessorter"
	"gopkg.in/yaml.v3"
)

const accessMask = 0o777

func manageOutputFileAndItsDirs(outputFilePath string, dir string, filename string) error {
	_, err := os.Stat(outputFilePath)
	if errors.Is(err, os.ErrNotExist) {
		err := filesmanager.MakeDirectory(dir)
		if err != nil {
			return err
		}

		err = filesmanager.CreateFile(dir, filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	configPathFlag := flag.String("config", "config.yaml", "path to config file")

	flag.Parse()

	configFile, err := config.GetConfigFile(configPathFlag)
	if err != nil {
		panic(err)
	}

	configData, err := config.GetConfigData(configFile)
	if err != nil {
		panic(err)
	}

	var files config.Config

	err = yaml.Unmarshal(configData, &files)
	if err != nil {
		panic(err)
	}

	dir, filename := filesmanager.ParseOutputFilePath(files.OutputFile)

	err = manageOutputFileAndItsDirs(files.OutputFile, dir, filename)
	if err != nil {
		panic(err)
	}

	inputFile, err := os.Open(files.InputFile)
	if err != nil && os.IsNotExist(err) {
		panic(err)
	}

	XMLDecoder := xml.NewDecoder(inputFile)
	XMLDecoder.CharsetReader = charsetsetter.Charset

	var CBCurrencyRate currencyrate.CurrencyRate

	if err := XMLDecoder.Decode(&CBCurrencyRate); err != nil {
		panic(err)
	}

	sort.Sort(valutessorter.ByValue(CBCurrencyRate))

	outputFile, err := os.OpenFile(path.Join(dir, filename), os.O_WRONLY, accessMask)
	if err != nil {
		panic(err)
	}

	JSONEncoder := json.NewEncoder(outputFile)
	JSONEncoder.SetIndent("", "\t")

	err = JSONEncoder.Encode(&CBCurrencyRate.Valutes)
	if err != nil {
		panic(err)
	}
}
