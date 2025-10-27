package jsonwriter

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func SaveToJSON(data any, filePath string) error {
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer func() {
		if err := outputFile.Close(); err != nil {
			panic(err)
		}
	}()

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")

	return encoder.Encode(data)
}
