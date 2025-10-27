package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	DirPerm  = 0o755
	FilePerm = 0o644
)

func WriteToJSON(data any, path string) error {
	err := os.MkdirAll(filepath.Dir(path), DirPerm)
	if err != nil {
		return fmt.Errorf("failed to create a dir: %w", err)
	}

	newdata, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("failed to write JSON: %w", err)
	}

	err = os.WriteFile(path, newdata, FilePerm)
	if err != nil {
		return fmt.Errorf("failed to write JSON: %w", err)
	}

	return nil
}
