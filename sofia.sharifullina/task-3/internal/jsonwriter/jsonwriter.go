package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const permission = 0o755

func SaveJSON(outputFile string, data any) error {
	dir := filepath.Dir(outputFile)

	err := os.MkdirAll(dir, permission)
	if err != nil {
		return fmt.Errorf("failed to creating directory: %w", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to creating file: %w", err)
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("failed to close file: %v\n", closeErr)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")

	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("failed to encoding JSON: %w", err)
	}

	return nil
}
