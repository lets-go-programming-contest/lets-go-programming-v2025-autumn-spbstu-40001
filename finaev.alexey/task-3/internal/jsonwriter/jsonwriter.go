package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func SaveJSON(outputPath string, data any) error {
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf("Failed create directory %s: %w", outputPath, err)
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("marshal JSON: %w", err)
	}
	if err := os.WriteFile(outputPath, jsonData, 0o644); err != nil {
		return fmt.Errorf("write file: %w", err)
	}
	return nil
}
