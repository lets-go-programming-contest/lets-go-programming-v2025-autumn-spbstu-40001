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
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("cannot create output directory: %w", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("cannot create output file: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("Warning: error closing file: %v\n", closeErr)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err = encoder.Encode(currencies); err != nil {
		return fmt.Errorf("cannot encode JSON data: %w", err)
	}

	return nil
}
