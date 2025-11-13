package jsonparser

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func WriteInfoFromCurrRateToOutputFile[T any](cbCurrencyRate *T, outputFilePath string, accessMask os.FileMode) error {
	dirAll := path.Dir(outputFilePath)

	if err := os.MkdirAll(dirAll, accessMask); err != nil {
		return fmt.Errorf("can't make directories of path %s: %w", outputFilePath, err)
	}

	outputFile, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_WRONLY, accessMask)
	if err != nil {
		return fmt.Errorf("can't open file %s: %w", path.Base(outputFilePath), err)
	}

	JSONEncoder := json.NewEncoder(outputFile)
	JSONEncoder.SetIndent("", "\t")

	if err := JSONEncoder.Encode(&cbCurrencyRate); err != nil {
		return fmt.Errorf("failed to encode currency rate to file %s: %w", outputFile.Name(), err)
	}

	if err = outputFile.Close(); err != nil {
		return fmt.Errorf("failed to close file %s: %w", outputFile.Name(), err)
	}

	return nil
}
