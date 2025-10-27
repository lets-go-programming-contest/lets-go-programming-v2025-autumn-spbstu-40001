package xmlparser

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/A1exMas1ov/task-3/internal/currency"
)

func ParseXML(path string) ([]currency.Valute, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read XML file: %w", err)
	}

	var valCurs currency.ValCurs

	err = xml.Unmarshal(data, &valCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	return valCurs.Valutes, nil
}
