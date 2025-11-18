package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func SaveToJSON(data any, filePath string) error {
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	outputFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	defer func() {
		if err := outputFile.Close(); err != nil {
			panic(err)
		}
	}()

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	return nil
}
