package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ParseJSON(path string, data interface{}, dirPerm, filePerm os.FileMode) error {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("serialize to JSON: %w", err)
	}

	directory := filepath.Dir(path)
	if err := os.MkdirAll(directory, dirPerm); err != nil {
		return fmt.Errorf("cant create directory '%s': %w", directory, err)
	}

	if err := os.WriteFile(path, jsonData, filePerm); err != nil {
		return fmt.Errorf("cant write to file '%s': %w", path, err)
	}

	return nil
}
