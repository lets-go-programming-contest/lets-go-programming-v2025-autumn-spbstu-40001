package jsonfile

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Save(path string, data any) error {
	err := os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "    ")
	err = enc.Encode(data)
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}
