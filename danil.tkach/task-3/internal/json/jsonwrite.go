package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Danil3352/task-3/internal/xml"
)

const (
	DirPerms  os.FileMode = 0o755
	FilePerms os.FileMode = 0o644
)

func WriteResult(currencies []xml.Currency, outputFile string) error {
	outputDir := filepath.Dir(outputFile)
	if err := os.MkdirAll(outputDir, DirPerms); err != nil {
		return fmt.Errorf("failed to create a dir %s: %w", outputDir, err)
	}

	jsonData, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to read to json: %w", err)
	}

	if err := os.WriteFile(outputFile, jsonData, FilePerms); err != nil {
		return fmt.Errorf("failed write file %s: %w", outputFile, err)
	}

	return nil
}
