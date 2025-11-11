package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func SaveJSON(path string, data any) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return fmt.Errorf("create dir for %s: %w", path, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file %s: %w", path, err)
	}

	defer func() {
		_ = file.Close()
	}()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "    ")

	if err := enc.Encode(data); err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	return nil
}
