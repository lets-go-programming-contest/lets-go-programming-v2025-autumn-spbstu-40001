package write

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func SerializeToJSON(path string, obj any, dirPermission, filePermission os.FileMode) error {
	data, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		return fmt.Errorf("json marshalling failed: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(path), dirPermission)
	if err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	err = os.WriteFile(path, data, filePermission)
	if err != nil {
		return fmt.Errorf(`failed to write file "%s": %w`, path, err)
	}

	return nil
}
