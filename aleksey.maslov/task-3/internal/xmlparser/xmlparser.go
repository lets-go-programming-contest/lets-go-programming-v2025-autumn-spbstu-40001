package xmlparser

import (
	"encoding/xml"
	"fmt"
	"os"
)

func ParseXML(path string, target any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read XML file: %w", err)
	}

	err = xml.Unmarshal(data, target)
	if err != nil {
		return fmt.Errorf("failed to parse XML: %w", err)
	}

	return nil
}
