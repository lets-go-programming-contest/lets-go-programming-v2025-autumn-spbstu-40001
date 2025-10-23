package files

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Aapng-cmd/task-3/internal/models"
	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

const (
	DirPerm  = 0o750 // rwxr-x---
	FilePerm = 0o600 // rw-------
)

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

type CommaFloat float64 // this one also has right for a living

func (cf *CommaFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// let it be here, please, i do not want one more packet in this small task
	var sIsAWorkingStringForFloatsWithCommaIsThisNameLongEnough string

	err := d.DecodeElement(&sIsAWorkingStringForFloatsWithCommaIsThisNameLongEnough, &start)
	if err != nil {
		return fmt.Errorf("ah, kozache, UnmarshalXML override func failed: %w", err)
	}

	sIsAWorkingStringForFloatsWithCommaIsThisNameLongEnough = strings.ReplaceAll(
		sIsAWorkingStringForFloatsWithCommaIsThisNameLongEnough,
		",",
		".",
	)

	val, err := strconv.ParseFloat(sIsAWorkingStringForFloatsWithCommaIsThisNameLongEnough, 64)
	if err != nil {
		return fmt.Errorf("ah, kozache, UnmarshalXML override func failed: %w", err)
	}

	*cf = CommaFloat(val)

	return nil
}

func ReadAndParseXML(xmlFilePath string) (models.ValCurs, error) {
	var valCurs models.ValCurs

	xmlData, err := os.ReadFile(xmlFilePath)
	if err != nil {
		return valCurs, fmt.Errorf("error reading XML file: %w", err)
	}

	// xmlData = []byte(strings.ReplaceAll(string(xmlData), ",", ".")) // i think this is less ram

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
	jsonData, err := json.MarshalIndent(valCurs.Valutes, "", "\t")
	if err != nil {
		return fmt.Errorf("WriteDataToJSON: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(jsonFilePath), DirPerm)
	if err != nil {
		return fmt.Errorf("WriteDataToJSON: %w", err)
	}

	err = os.WriteFile(jsonFilePath, jsonData, FilePerm)
	if err != nil {
		return fmt.Errorf("WriteDataToJSON: %w", err)
	}

	return nil
}
