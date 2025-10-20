package jsonwriter

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func SaveToJSON(data any, filePath string) {
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		panic(err)
	}

	output_file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer output_file.Close()

	encoder := json.NewEncoder(output_file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		panic(err)
	}
}
