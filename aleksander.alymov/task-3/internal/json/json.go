package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	dirPerms = 0o755
)

func WriteJSON(data interface{}, outputPath string) error {
	if err := os.MkdirAll(filepath.Dir(outputPath), dirPerms); err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	return nil
}
