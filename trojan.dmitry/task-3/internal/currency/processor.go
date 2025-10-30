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

	if vc == nil || len(vc.Valutes) == 0 {
		return fmt.Errorf("no currency data found in XML file")
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
			continue
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
	if len(results) == 0 {
		return fmt.Errorf("no valid currencies processed after filtering")
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Value != results[j].Value {
			return results[i].Value > results[j].Value
		}
		if results[i].NumCode != results[j].NumCode {
			return results[i].NumCode < results[j].NumCode
		}
		return results[i].CharCode < results[j].CharCode
	})

	err = jsonfile.Save(outputPath, results)
	if err != nil {
		return fmt.Errorf("save json: %w", err)
	}
	return nil
}
