package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	dirPermissions  = 0o755
	filePermissions = 0o600
)

func ParseJSON(filePath string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("serialize to JSON: %w", err)
	}

	directory := filepath.Dir(filePath)

	if err := os.MkdirAll(directory, dirPermissions); err != nil {
		return fmt.Errorf("cannot create directory '%s': %w", directory, err)
	}

	if err := os.WriteFile(filePath, jsonData, filePermissions); err != nil {
		return fmt.Errorf("cannot write to file '%s': %w", filePath, err)
	}

	return nil
}
