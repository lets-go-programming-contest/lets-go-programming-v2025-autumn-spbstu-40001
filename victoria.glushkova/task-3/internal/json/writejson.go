package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func WriteCurrencyRateToFile[T any](data *T, outputFilePath string, accessMask os.FileMode) error {
	dirAll := filepath.Dir(outputFilePath)

	if err := os.MkdirAll(dirAll, accessMask); err != nil {
		return fmt.Errorf("can't make directories of path %s: %w", outputFilePath, err)
	}

	outputFile, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_WRONLY, accessMask)
	if err != nil {
		return fmt.Errorf("can't open file %s: %w", filepath.Base(outputFilePath), err)
	}

	defer func() {
		err := outputFile.Close()
		panicIfErr(err)
	}()

	JSONEncoder := json.NewEncoder(outputFile)
	JSONEncoder.SetIndent("", "\t")

	if err := JSONEncoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode currency rate to file %s: %w", outputFile.Name(), err)
	}

	return nil
}
