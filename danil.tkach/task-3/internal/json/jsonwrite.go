package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func WriteResult[T any](data T, outputFile string, dirPerms, filePerms os.FileMode) error {
	outputDir := filepath.Dir(outputFile)
	if err := os.MkdirAll(outputDir, dirPerms); err != nil {
		return fmt.Errorf("failed to create a dir %s: %w", outputDir, err)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data to json: %w", err)
	}

	if err := os.WriteFile(outputFile, jsonData, filePerms); err != nil {
		return fmt.Errorf("failed write file %s: %w", outputFile, err)
	}

	return nil
}
