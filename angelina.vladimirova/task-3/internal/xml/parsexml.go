package xml

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

func ParseXML(filePath string, data interface{}) error {
	xmlData, err := xml.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("serialize to XML: %w", err)
	}

	directory := filepath.Dir(filePath)
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return fmt.Errorf("cannot create directory '%s': %w", directory, err)
	}

	if err := os.WriteFile(filePath, xmlData, 0o600); err != nil {
		return fmt.Errorf("cannot write to file '%s': %w", filePath, err)
	}

	return nil
}
