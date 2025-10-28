package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/A1exMas1ov/task-3/internal/currency"
)

func SaveJSON(path string, valutes []currency.Valute) error {

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")

	if err := encoder.Encode(valutes); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil

}
