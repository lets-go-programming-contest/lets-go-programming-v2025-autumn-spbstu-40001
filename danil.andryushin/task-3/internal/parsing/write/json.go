package write

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ToJSON(path string, obj any, permission os.FileMode) error {
	data, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return fmt.Errorf("json marshalling failed: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(path), permission)
	if err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	err = os.WriteFile(path, data, permission)
	if err != nil {
		return fmt.Errorf(`failed to write file "%s": %w`, path, err)
	}

	return nil
}
