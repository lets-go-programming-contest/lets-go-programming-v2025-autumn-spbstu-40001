package processor

import (
	"fmt"
	"sort"

	"github.com/netwite/task-3/internal/currency"
	"github.com/netwite/task-3/internal/json"
	"github.com/netwite/task-3/internal/xml"
)

type Processor struct{}

func NewProcessor() *Processor {
	return &Processor{}
}

func (p *Processor) Process(inputFile, outputFile string) error {
	var valCurs currency.ValCurs

	if err := xml.DecodeXMLFile(inputFile, &valCurs); err != nil {
		return fmt.Errorf("failed to read and parse XML: %w", err)
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	if err := json.WriteResult(valCurs.Valutes, outputFile); err != nil {
		return fmt.Errorf("failed to write JSON result: %w", err)
	}

	return nil
}
