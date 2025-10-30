package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func WriteToFile(outputFile string, data interface{}) error {
	outputDir := filepath.Dir(outputFile)

	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		return fmt.Errorf("cannot create output directory: %w", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("cannot create output file: %w", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	err = encoder.Encode(data)
	if err != nil {
		_ = file.Close()

		return fmt.Errorf("cannot encode JSON data: %w", err)
	}

	err = file.Close()
	if err != nil {
		return fmt.Errorf("cannot close output file: %w", err)
	}

	return nil
}
