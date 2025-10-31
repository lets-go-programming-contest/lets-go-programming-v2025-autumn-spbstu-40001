package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	DirPerm  os.FileMode = 0o755
	FilePerm os.FileMode = 0o644
)

func SaveJSON(outputPath string, data any) error {
	return SaveJSONwithPerms(outputPath, data, DirPerm, FilePerm)
}

func SaveJSONwithPerms(outputPath string, data any, dirPerm, filePerm os.FileMode) error {
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
