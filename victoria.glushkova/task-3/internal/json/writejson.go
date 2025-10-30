package json

import (
	"encoding/json"
	"fmt"
	"github.com/vikaglushkova/task-3/internal/currency"
	"os"
	"path/filepath"
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
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	err = encoder.Encode(currencies)
	if err != nil {
		return fmt.Errorf("cannot encode JSON data: %w", err)
	}

	return nil
}
