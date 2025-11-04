package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/Kirill2155/task-3/internal/model"
	"golang.org/x/net/html/charset"
)

func ParserXML(path string) (*model.ValCurs, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("XML file does not exist: %w", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read XML file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs model.ValCurs

	err = decoder.Decode(&valCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	return &valCurs, nil
}
