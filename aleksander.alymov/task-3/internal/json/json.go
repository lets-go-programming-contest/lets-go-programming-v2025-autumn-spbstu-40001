package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type FileSaver interface {
	Save(filename string, data interface{}) error
}

type JSONSaver struct {
	DirPerms  os.FileMode
	FilePerms os.FileMode
}

func NewSaver() *JSONSaver {
	return &JSONSaver{
		DirPerms:  0755,
		FilePerms: 0644,
	}
}

func (s *JSONSaver) Save(filename string, data interface{}) error {
	outputDir := filepath.Dir(filename)
	if err := os.MkdirAll(outputDir, s.DirPerms); err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal JSON: %w", err)
	}

	if err := os.WriteFile(filename, jsonData, s.FilePerms); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
