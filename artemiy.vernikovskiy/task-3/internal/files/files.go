// Package files provides utilities for reading configuration files, parsing XML data, and writing JSON output.
// This package now includes generic functions for universal parsing and writing of data.
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

// ReadYAML reads and parses a YAML file into a generic type T.
// It returns the parsed data or an error.
func ReadYAML[T any](yamlPath string) (T, error) {
	var data T

	fileData, err := os.ReadFile(yamlPath)
	if err != nil {
		return data, fmt.Errorf("error reading YAML file: %w", err)
	}

	err = yaml.Unmarshal(fileData, &data)
	if err != nil {
		return data, fmt.Errorf("error unmarshaling YAML: %w", err)
	}

	return data, nil
}

// ReadYAMLConfigFile reads and parses a YAML configuration file.
// It returns the input file path, output file path, and any error encountered.
// This function is now implemented using the generic ReadYAML function for better reusability.
func ReadYAMLConfigFile(yamlPath string) (string, string, error) {
	settings, err := ReadYAML[models.Settings](yamlPath)
	if err != nil {
		return "", "", err
	}

	return settings.InputFileSetting, settings.OutputFileSetting, nil
}

// ReadXML reads and parses an XML file into a generic type T.
// It handles character encoding and returns the parsed data or an error.
func ReadXML[T any](xmlFilePath string) (T, error) {
	var data T

	xmlData, err := os.ReadFile(xmlFilePath)
	if err != nil {
		return data, fmt.Errorf("error reading XML file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, encoding)
	}

	err = decoder.Decode(&data)
	if err != nil {
		return data, fmt.Errorf("error unmarshaling XML: %w", err)
	}

	return data, nil
}

// ReadAndParseXML reads an XML file and parses it into a ValCurs structure.
// It handles character encoding and returns the parsed data or an error.
// This function is now implemented using the generic ReadXML function for better reusability.
func ReadAndParseXML(xmlFilePath string) (models.ValCurs, error) {
	return ReadXML[models.ValCurs](xmlFilePath)
}

// WriteJSON writes the provided data of generic type T to a JSON file.
// It creates necessary directories and returns any error encountered.
func WriteJSON[T any](data T, jsonFilePath string) error {
	jsonData, err := json.MarshalIndent(data, "", "\t")
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

// WriteDataToJSON writes the ValCurs data to a JSON file.
// It creates necessary directories and returns any error encountered.
// This function is now implemented using the generic WriteJSON function for better reusability.
func WriteDataToJSON(valCurs models.ValCurs, jsonFilePath string) error {
	return WriteJSON(valCurs.Valutes, jsonFilePath)
}
