package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"sort"

	"github.com/Nekich06/task-3/internal/config"
	"github.com/Nekich06/task-3/internal/currencyrate"
	"github.com/Nekich06/task-3/internal/filesmanager"
	"github.com/Nekich06/task-3/internal/valutessorter"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
	"gopkg.in/yaml.v3"
)

const accessMask = 0o777

func getConfigInfo(configPath *string) (config.Config, error) {
	var configInfo config.Config

	configFile, err := config.GetConfigFile(configPath)
	if err != nil {
		return configInfo, fmt.Errorf("can't get config file descriptor: %w", err)
	}

	configData, err := config.GetConfigData(configFile)
	if err != nil {
		return configInfo, fmt.Errorf("can't get config data: %w", err)
	}

	err = yaml.Unmarshal(configData, &configInfo)
	if err != nil {
		return configInfo, fmt.Errorf("can't unmarshal config data: %w", err)
	}

	return configInfo, nil
}

func manageOutputFileAndItsDirs(outputFilePath string, dir string, filename string) error {
	_, err := os.Stat(outputFilePath)
	if errors.Is(err, os.ErrNotExist) {
		err := filesmanager.MakeDirectory(dir)
		if err != nil {
			return fmt.Errorf("in main: %w", err)
		}

		err = filesmanager.CreateFile(dir, filename)
		if err != nil {
			return fmt.Errorf("in main: %w", err)
		}
	}

	return nil
}

func writeInfoFromInputFileToCurrRate(inputFile *os.File, cbCurrencyRate *currencyrate.CurrencyRate) error {
	XMLDecoder := xml.NewDecoder(inputFile)
	XMLDecoder.CharsetReader = charset.NewReader

	if err := XMLDecoder.Decode(&cbCurrencyRate); err != nil {
		return fmt.Errorf("failed to decode file %s: %w", inputFile.Name(), err)
	}

	return nil
}

func writeInfoFromCurrRateToOutputFile(cbCurrencyRate *currencyrate.CurrencyRate, outputFile *os.File) error {
	JSONEncoder := json.NewEncoder(outputFile)
	JSONEncoder.SetIndent("", "\t")

	if err := JSONEncoder.Encode(&cbCurrencyRate.Valutes); err != nil {
		return fmt.Errorf("failed to encode currency rate to file %s: %w", outputFile.Name(), err)
	}

	return nil
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")

	flag.Parse()

	configInfo, err := getConfigInfo(configPath)
	if err != nil {
		panic(err)
	}

	dirPath, filename := filesmanager.ParseOutputFilePath(configInfo.OutputFile)

	err = manageOutputFileAndItsDirs(configInfo.OutputFile, dirPath, filename)
	if err != nil {
		panic(err)
	}

	inputFile, err := os.Open(configInfo.InputFile)
	if err != nil && os.IsNotExist(err) {
		panic(err)
	}

	var CBCurrencyRate currencyrate.CurrencyRate

	err = writeInfoFromInputFileToCurrRate(inputFile, &CBCurrencyRate)
	if err != nil {
		panic(err)
	}

	sort.Sort(valutessorter.ByValue(CBCurrencyRate))

	outputFile, err := os.OpenFile(path.Join(dirPath, filename), os.O_WRONLY, accessMask)
	if err != nil {
		panic(err)
	}

	err = writeInfoFromCurrRateToOutputFile(&CBCurrencyRate, outputFile)
	if err != nil {
		panic(err)
	}
}
