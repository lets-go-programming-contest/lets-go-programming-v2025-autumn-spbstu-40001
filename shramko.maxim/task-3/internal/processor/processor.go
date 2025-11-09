package processor

import (
	"fmt"

	"github.com/Elektrek/task-3/internal/config"
	"github.com/Elektrek/task-3/internal/parser"
	"github.com/Elektrek/task-3/internal/sorter"
	"github.com/Elektrek/task-3/internal/writer"
)

func ProcessCurrencies(cfg *config.Config) error {
	collection, err := parser.ParseCurrencies(cfg.InputFile)
	if err != nil {
		return fmt.Errorf("failed to parse currencies: %w", err)
	}

	sorter.SortByValueDescending(collection.Currencies)

	if err := writer.WriteJSON(cfg.OutputFile, collection.Currencies); err != nil {
		return fmt.Errorf("failed to write JSON: %w", err)
	}

	return nil
}
