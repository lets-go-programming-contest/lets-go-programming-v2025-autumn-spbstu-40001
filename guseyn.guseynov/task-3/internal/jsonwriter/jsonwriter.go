package jsonwriter

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Write(path string, data any, dirPerm os.FileMode, filePerm os.FileMode) {
	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, dirPerm)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")

	err = encoder.Encode(data)
	if err != nil {
		panic(err)
	}
}
