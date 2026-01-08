package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func WriteToJSON(data any, path string, dirPerm os.FileMode, filePerm os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(path), dirPerm)
	if err != nil {
		return fmt.Errorf("failed to create a dir: %w", err)
	}

	newdata, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("failed to write JSON: %w", err)
	}

	err = os.WriteFile(path, newdata, filePerm)
	if err != nil {
		return fmt.Errorf("failed to write JSON: %w", err)
	}

	return nil
}
