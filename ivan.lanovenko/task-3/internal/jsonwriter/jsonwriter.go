package jsonwriter

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func SaveToJSON(data any, filePath string) {
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		panic(err)
	}

	outputFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := outputFile.Close(); err != nil {
			panic(err)
		}
	}()

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(data); err != nil {
		panic(err)
	}
}
