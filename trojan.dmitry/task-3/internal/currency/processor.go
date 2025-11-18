package currency

import (
	"fmt"
	"sort"

	"github.com/DimasFantomasA/task-3/internal/cbrusxml"
	"github.com/DimasFantomasA/task-3/internal/jsonfile"
)

func Process(inputPath, outputPath string) error {
	valCurs, err := cbrusxml.ParseFile(inputPath)
	if err != nil {
		return fmt.Errorf("parse xml: %w", err)
	}

	valutes := transform(valCurs)
	sortValutes(valutes)

	err = jsonfile.Save(outputPath, valutes)
	if err != nil {
		return fmt.Errorf("save json: %w", err)
	}

	return nil
}

func transform(valCurs *cbrusxml.ValCurs) []cbrusxml.Valute {
	return append([]cbrusxml.Valute{}, valCurs.Valutes...)
}

func sortValutes(valutes []cbrusxml.Valute) {
	sort.Slice(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})
}
