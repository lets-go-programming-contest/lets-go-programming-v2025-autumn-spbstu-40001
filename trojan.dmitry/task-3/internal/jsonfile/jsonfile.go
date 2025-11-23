package jsonfile

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func Save(path string, data any, perm fs.FileMode) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("mkdir %s: %w", dir, err)
	}

	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	tmpFile, err := os.CreateTemp(dir, "tmp-*.json")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}

	defer func() {
		if cerr := tmpFile.Close(); cerr != nil {
			panic(fmt.Errorf("close temp file: %w", cerr))
		}
	}()

	if _, err := tmpFile.Write(content); err != nil {
		return fmt.Errorf("write temp file: %w", err)
	}

	if err := os.Rename(tmpFile.Name(), path); err != nil {
		return fmt.Errorf("rename temp file: %w", err)
	}

	if err := os.Chmod(path, perm); err != nil {
		return fmt.Errorf("chmod: %w", err)
	}

	return nil
}
