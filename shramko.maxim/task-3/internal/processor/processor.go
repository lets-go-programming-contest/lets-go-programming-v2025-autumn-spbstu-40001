package processor

import (
	"fmt"

	"github.com/Elektrek/task-3/internal/config"
	"github.com/Elektrek/task-3/internal/parser"
	"github.com/Elektrek/task-3/internal/sorter"
	"github.com/Elektrek/task-3/internal/writer"
)

func ProcessCurrencies(cfg *config.Config) error {
	currencyCollection, parseErr := parser.ParseCurrencies(cfg.InputFile)
	if parseErr != nil {
		return fmt.Errorf("failed to parse currencies: %w", parseErr)
	}

	sorter.SortByValueDescending(currencyCollection.CurrencyItems)

	if writeErr := writer.WriteJSON(cfg.OutputFile, currencyCollection.CurrencyItems); writeErr != nil {
		return fmt.Errorf("failed to write JSON: %w", writeErr)
	}

	return nil
}
