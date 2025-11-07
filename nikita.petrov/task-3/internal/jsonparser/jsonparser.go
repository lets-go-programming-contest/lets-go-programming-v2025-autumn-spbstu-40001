package jsonparser

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/Nekich06/task-3/internal/currencyrate"
	"github.com/Nekich06/task-3/internal/filesmanager"
)

const accessMask = 0o777

func WriteInfoFromCurrRateToOutputFile(cbCurrencyRate *currencyrate.CurrencyRate, outputFilePath string) error {
	outputFile, err := os.OpenFile(outputFilePath, os.O_WRONLY, accessMask)
	if err != nil {
		return fmt.Errorf("can't open file %s: %w", path.Base(outputFilePath), err)
	}

	dir, filename := filesmanager.ParseOutputFilePath(outputFilePath)

	_, err = os.Stat(outputFilePath)
	if errors.Is(err, os.ErrNotExist) {
		err := filesmanager.MakeDirectory(dir)
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		err = filesmanager.CreateFile(dir, filename)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	JSONEncoder := json.NewEncoder(outputFile)
	JSONEncoder.SetIndent("", "\t")

	if err := JSONEncoder.Encode(&cbCurrencyRate.Valutes); err != nil {
		return fmt.Errorf("failed to encode currency rate to file %s: %w", outputFile.Name(), err)
	}

	return nil
}
