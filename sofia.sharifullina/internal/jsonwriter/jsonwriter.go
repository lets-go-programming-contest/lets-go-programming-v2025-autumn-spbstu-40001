package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func SaveJSON(outputFile string, data any,  filePermission, dirPermission os.FileMode) error {
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
		err = file.Close()
		if err != nil {
			panic("failed to close file in saveJSON")
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")

	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("failed to encoding JSON: %w", err)
	}

	return nil
}
