package funcs

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Aapng-cmd/task-3/internal/models"
	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

func SortDataByValue(valCurs models.ValCurs) models.ValCurs {
	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	return valCurs
}

func ReadYAMLConfigFile(yamlPath string) (string, string, error) {
	var settings models.Settings

	data, err := os.ReadFile(yamlPath)
	if err != nil {
		return "", "", fmt.Errorf("error reading YAML config file: %w", err)
	}

	err = yaml.Unmarshal(data, &settings)
	if err != nil {
		return "", "", fmt.Errorf("error unmarshaling YAML: %w", err)
	}

	return settings.InputFileSetting, settings.OutputFileSetting, nil
}

func ReadAndParseXML(xmlFilePath string) (models.ValCurs, error) {
	var valCurs models.ValCurs

	xmlData, err := os.ReadFile(xmlFilePath)
	if err != nil {
		return valCurs, fmt.Errorf("error reading XML file: %w", err)
	}

	xmlString := string(xmlData)
	xmlString = strings.ReplaceAll(xmlString, ",", ".") // F*cking xml with comma separation
	xmlData = []byte(xmlString)

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, encoding)
	}

	err = decoder.Decode(&valCurs)
	if err != nil {
		return valCurs, fmt.Errorf("error unmarshaling XML: %w", err)
	}

	return valCurs, nil
}

func WriteDataToJSON(valCurs models.ValCurs, jsonFilePath string) error {
	valCurs = SortDataByValue(valCurs)

	jsonData, err := json.MarshalIndent(valCurs.Valutes, "", "\t")
	if err != nil {
		return fmt.Errorf("WriteDataToJSON: %w", err)
	}

	dir := filepath.Dir(jsonFilePath)

	err = os.MkdirAll(dir, 0o750)
	if err != nil {
		return fmt.Errorf("WriteDataToJSON: %w", err)
	}

	err = os.WriteFile(jsonFilePath, jsonData, 0o600)
	if err != nil {
		return fmt.Errorf("WriteDataToJSON: %w", err)
	}

	return nil
}
