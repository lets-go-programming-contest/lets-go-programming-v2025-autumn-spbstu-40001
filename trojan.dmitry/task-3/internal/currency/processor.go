package currency

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/DimasFantomasA/task-3/internal/cbrusxml"
	"github.com/DimasFantomasA/task-3/internal/jsonfile"
)

type Result struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func Process(inputPath, outputPath string) error {
	vc, err := cbrusxml.ParseFile(inputPath)
	if err != nil {
		return fmt.Errorf("parse xml: %w", err)
	}

	results := make([]Result, 0, len(vc.Valutes))
	for _, v := range vc.Valutes {
		valStr := strings.TrimSpace(v.Value)
		if valStr == "" {
			continue
		}

		valStr = strings.ReplaceAll(valStr, ",", ".")
		parsed, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return fmt.Errorf("Error:", v.CharCode, err)
		}

		nominal := float64(v.Nominal)
		if nominal == 0 {
			nominal = 1
		}
		valuePerOne := parsed / nominal

		results = append(results, Result{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    valuePerOne,
		})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Value > results[j].Value
	})

	err = jsonfile.Save(outputPath, results)
	if err != nil {
		return fmt.Errorf("Error:", err)
	}
	return nil
}
