package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func WriteToFile(outputFile string, data interface{}, dirPermissions os.FileMode) error {
	outputDir := filepath.Dir(outputFile)

	err := os.MkdirAll(outputDir, dirPermissions)
	if err != nil {
		return fmt.Errorf("cannot create output directory: %w", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("cannot create output file: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(fmt.Sprintf("cannot close output file: %v", err))
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("cannot encode JSON data: %w", err)
	}

	return nil
}
