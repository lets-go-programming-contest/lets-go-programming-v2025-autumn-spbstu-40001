package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/vikaglushkova/task-3/internal/currency"
)

func WriteToFile(outputFile string, currencies []currency.Currency) error {
	outputDir := filepath.Dir(outputFile)
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		return fmt.Errorf("cannot create output directory: %w", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("cannot create output file: %w", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err = encoder.Encode(currencies); err != nil {
		file.Close()
		return fmt.Errorf("cannot encode JSON data: %w", err)
	}

	file.Close()
	return nil
}
