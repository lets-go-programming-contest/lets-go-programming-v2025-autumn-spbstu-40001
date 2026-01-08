package currency

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func (list outputList) WriteJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(list); err != nil {
		return fmt.Errorf("json encode failed: %w", err)
	}

	return nil
}

func (list outputList) WriteJSONFile(path string) error {
	dir := filepath.Dir(path)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("create directory %s: %w", dir, err)
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file %s: %w", path, err)
	}
	defer file.Close()

	return list.WriteJSON(file)
}
