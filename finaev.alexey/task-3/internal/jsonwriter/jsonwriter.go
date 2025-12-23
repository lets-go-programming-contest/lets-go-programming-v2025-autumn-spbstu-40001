package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func SaveJSON(outputPath string, data any, dirPerm, filePerm os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(outputPath), dirPerm); err != nil {
		return fmt.Errorf("failed create directory %s: %w", outputPath, err)
	}

	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("marshal JSON: %w", err)
	}

	err = os.WriteFile(outputPath, jsonData, filePerm)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
