// Package files provides utilities for reading configuration files, parsing XML data, and writing JSON output.
package files

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Aapng-cmd/task-3/internal/models"
	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

const (
	// DirPerm defines the permissions for created directories.
	DirPerm = 0o750 // rwxr-x---
	// FilePerm defines the permissions for created files.
	FilePerm = 0o600 // rw-------
)

// ReadYAMLConfigFile reads and parses a YAML configuration file.
// It returns the input file path, output file path, and any error encountered.
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

// ReadAndParseXML reads an XML file and parses it into a ValCurs structure.
// It handles character encoding and returns the parsed data or an error.
func ReadAndParseXML(xmlFilePath string) (models.ValCurs, error) {
	var valCurs models.ValCurs

	xmlData, err := os.ReadFile(xmlFilePath)
	if err != nil {
		return valCurs, fmt.Errorf("error reading XML file: %w", err)
	}

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

// WriteDataToJSON writes the ValCurs data to a JSON file.
// It creates necessary directories and returns any error encountered.
func WriteDataToJSON(valCurs models.ValCurs, jsonFilePath string) error {
	jsonData, err := json.MarshalIndent(valCurs.Valutes, "", "\t")
	if err != nil {
		return fmt.Errorf("error marshaling to JSON: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(jsonFilePath), DirPerm)
	if err != nil {
		return fmt.Errorf("error creating directories: %w", err)
	}

	err = os.WriteFile(jsonFilePath, jsonData, FilePerm)
	if err != nil {
		return fmt.Errorf("error writing JSON file: %w", err)
	}

	return nil
}
