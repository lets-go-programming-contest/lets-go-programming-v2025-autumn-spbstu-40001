package currency

import (
	"fmt"
	"sort"

	"github.com/DimasFantomasA/task-3/internal/cbrusxml"
	"github.com/DimasFantomasA/task-3/internal/jsonfile"
	"github.com/DimasFantomasA/task-3/internal/xmlparser"
)

func Process(inputPath, outputPath string) error {
	valCurs, err := LoadValutes(inputPath)
	if err != nil {
		return err
	}

	valutes := PrepareValutes(valCurs)
	SortValutes(valutes)

	err = StoreValutes(outputPath, valutes)
	if err != nil {
		return err
	}

	return nil
}

func LoadValutes(path string) (*cbrusxml.ValCurs, error) {
	valCurs, err := xmlparser.ParseXMLFile[cbrusxml.ValCurs](path)
	if err != nil {
		return nil, fmt.Errorf("parse xml: %w", err)
	}
	return valCurs, nil
}

func PrepareValutes(valCurs *cbrusxml.ValCurs) []cbrusxml.Valute {
	return append([]cbrusxml.Valute{}, valCurs.Valutes...)
}

func SortValutes(valutes []cbrusxml.Valute) {
	sort.Slice(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})
}

func StoreValutes(path string, valutes []cbrusxml.Valute) error {
	if err := jsonfile.Save(path, valutes, 0o755); err != nil {
		return fmt.Errorf("save json: %w", err)
	}

	return nil
}
