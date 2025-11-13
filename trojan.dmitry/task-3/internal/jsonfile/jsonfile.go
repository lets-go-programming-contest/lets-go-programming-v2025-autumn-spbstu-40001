package jsonfile

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const defaultDirPerm = 0o755

func Save(path string, data any) error {
	err := os.MkdirAll(filepath.Dir(path), defaultDirPerm)
	if err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil && err == nil {
			err = fmt.Errorf("close file %w", closeErr)
		}
	}()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "    ")

	err = enc.Encode(data)
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}
