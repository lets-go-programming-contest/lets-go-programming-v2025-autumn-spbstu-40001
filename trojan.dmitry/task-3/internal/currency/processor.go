package currency

import (
	"fmt"
	"sort"

	"github.com/DimasFantomasA/task-3/internal/cbrusxml"
	"github.com/DimasFantomasA/task-3/internal/jsonfile"
)

type Result struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func Process(inputPath, outputPath string) error {
	valCurs, err := cbrusxml.ParseFile(inputPath)

	if err != nil {
		return fmt.Errorf("parse xml: %w", err)
	}

	results := make([]Result, 0, len(valCurs.Valutes))

	for _, val := range valCurs.Valutes {
		results = append(results, Result{
			NumCode:  val.NumCode,
			CharCode: val.CharCode,
			Value:    float64(val.Value),
		})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Value > results[j].Value
	})

	err = jsonfile.Save(outputPath, results)
	if err != nil {
		return fmt.Errorf("save json: %w", err)
	}

	return nil
}
