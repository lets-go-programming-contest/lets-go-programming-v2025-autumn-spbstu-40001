package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func WriteFile[T any](path string, data T, defaultMode os.FileMode) error {
	serialized, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("failed to serialize data to json: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(path), defaultMode)
	if err != nil {
		return fmt.Errorf("cannot create required directories: %w", err)
	}

	err = os.WriteFile(path, serialized, defaultMode)
	if err != nil {
		return fmt.Errorf("cannot write output file: %w", err)
	}

	return nil
}
