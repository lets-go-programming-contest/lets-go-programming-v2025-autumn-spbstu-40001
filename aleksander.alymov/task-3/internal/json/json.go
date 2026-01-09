package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	DirPerms  os.FileMode = 0o755
	FilePerms os.FileMode = 0o644
)

func WriteResult[T any](data T, outputFile string) error {
	outputDir := filepath.Dir(outputFile)
	if err := os.MkdirAll(outputDir, DirPerms); err != nil {
		return fmt.Errorf("failed to create a dir %s: %w", outputDir, err)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data to json: %w", err)
	}

	if err := os.WriteFile(outputFile, jsonData, FilePerms); err != nil {
		return fmt.Errorf("failed write file %s: %w", outputFile, err)
	}

	return nil
}
