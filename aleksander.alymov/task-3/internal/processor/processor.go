package processor

import (
	"fmt"

	"github.com/netwite/task-3/internal/currency"
	"github.com/netwite/task-3/internal/json"
	"github.com/netwite/task-3/internal/sorter"
	"github.com/netwite/task-3/internal/xml"
)

type DataProcessor struct {
	loader    xml.FileLoader
	converter currency.Converter
	sorter    sorter.Sorter
	saver     json.FileSaver
}

func NewDataProcessor(loader xml.FileLoader, converter currency.Converter, sorter sorter.Sorter, saver json.FileSaver) *DataProcessor { // Измените тип параметра
	return &DataProcessor{
		loader:    loader,
		converter: converter,
		sorter:    sorter,
		saver:     saver,
	}
}

func (p *DataProcessor) Process(inputFile, outputFile string) error {
	var sourceData currency.XMLValCurs // ← используем конкретный тип

	if err := p.loader.Load(inputFile, &sourceData); err != nil {
		return fmt.Errorf("load data: %w", err)
	}

	convertedData, err := p.converter.Convert(&sourceData) // ← передаем указатель
	if err != nil {
		return fmt.Errorf("convert data: %w", err)
	}

	if sortable, ok := convertedData.(sorter.Sortable); ok && p.sorter != nil {
		p.sorter.Sort(sortable)
	}

	if err := p.saver.Save(outputFile, convertedData); err != nil {
		return fmt.Errorf("save data: %w", err)
	}

	return nil
}
