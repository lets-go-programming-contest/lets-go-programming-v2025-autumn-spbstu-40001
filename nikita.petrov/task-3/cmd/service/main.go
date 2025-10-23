package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"os"
	"path"
	"sort"

	"github.com/Nekich06/task-3/internal/currencyrate"
	"github.com/Nekich06/task-3/internal/fmanager"
	"github.com/Nekich06/task-3/internal/valutessorter"

	"gopkg.in/yaml.v3"
)

func main() {
	var configPathFlag = flag.String("config", "config.yaml", "path to config file")

	flag.Parse()

	configFile := fmanager.GetConfigFile(configPathFlag)
	configData := fmanager.GetConfigData(configFile)

	var files fmanager.Config
	err := yaml.Unmarshal(configData, &files)

	if err != nil {
		panic(err)
	}

	dir, filename := fmanager.ParseOutputFilePath(files.OutputFile)

	_, err = os.Stat(files.OutputFile)
	if errors.Is(err, os.ErrNotExist) {
		fmanager.MakeDirectory(dir)
		fmanager.CreateFile(dir, filename)
	}

	inputFile, err := os.Open(files.InputFile)
	if err != nil && os.IsNotExist(err) {
		panic(err)
	}

	XMLDecoder := xml.NewDecoder(inputFile)
	XMLDecoder.CharsetReader = fmanager.Charset

	var CBCurrencyRate currencyrate.CurrencyRate

	if err := XMLDecoder.Decode(&CBCurrencyRate); err != nil {
		panic(err)
	}

	sort.Sort(valutessorter.ByValue(CBCurrencyRate))

	outputFile, err := os.OpenFile(path.Join(dir, filename), os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	JSONEncoder := json.NewEncoder(outputFile)
	JSONEncoder.SetIndent("", "\t")
	err = JSONEncoder.Encode(&CBCurrencyRate.Valute)

	if err != nil {
		panic(err)
	}
}
