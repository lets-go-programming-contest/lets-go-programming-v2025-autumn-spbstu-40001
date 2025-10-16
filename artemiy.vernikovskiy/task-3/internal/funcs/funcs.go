package funcs

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/Aapng-cmd/task-3/internal/models"
	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

func ReadYAMLConfigFile(yamlPath string) (string, string, error) {
	var settings models.Settings

	data, err := os.ReadFile(yamlPath)
	if err != nil {
		fmt.Println("Error reading YAML config file: %v", err)

		return "", "", err
	}

	err = yaml.Unmarshal(data, &settings)
	if err != nil {
		fmt.Println("Error unmarshaling YAML: %v", err)

		return "", "", err
	}

	return settings.InputFileSetting, settings.OutputFileSetting, nil
}

func ReadAndParseXML(xmlFilePath string) (models.ValCurs, error) {
	var valCurs models.ValCurs

	xmlData, err := os.ReadFile(xmlFilePath)
	if err != nil {
		fmt.Println("Error reading XML file: %v", err)

		return valCurs, err
	}

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, encoding)
	}

	err = decoder.Decode(&valCurs)
	if err != nil {
		fmt.Println("Error unmarshaling XML: %v", err)

		return valCurs, err
	}

	return valCurs, nil
}

func WriteDataToJSON(valCurs models.ValCurs, JSONFilePath string) error {
	jsonData, err := json.MarshalIndent(valCurs.Valutes, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(JSONFilePath, jsonData, 0o644)
	if err != nil {
		return err
	}

	return nil
}
