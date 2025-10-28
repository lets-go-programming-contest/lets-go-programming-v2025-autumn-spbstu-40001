package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/Danil3352/task-3/internal/models"
	"golang.org/x/net/html/charset"
)

func ReadValCurs(inputFile string) (*models.ValCurs, error) {
	xmlData, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read XML-file %s: %w", inputFile, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))

	decoder.CharsetReader = func(c string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, c)
	}

	var valCurs models.ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("failed to parse XML-file: %w", err)
	}

	return &valCurs, nil
}
