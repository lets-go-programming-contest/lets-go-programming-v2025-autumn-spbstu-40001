package currency

import (
	"fmt"
	"os"
	"sort"

	"github.com/Danil3352/task-3/internal/json"
	"github.com/Danil3352/task-3/internal/models"
	"github.com/Danil3352/task-3/internal/xml"
)

type ByValue []models.Currency

func (a ByValue) Len() int {
	return len(a)
}

func (a ByValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByValue) Less(i, j int) bool {
	return a[i].Value > a[j].Value
}

func Process(inputFile, outputFile string) error {
	valCurs, err := xml.ReadValCurs(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read and parse XML: %w", err)
	}

	sort.Sort(ByValue(valCurs.Valutes))

	const (
		DirPerms  os.FileMode = 0o755
		FilePerms os.FileMode = 0o644
	)

	if err := json.WriteResult(valCurs.Valutes, outputFile, DirPerms, FilePerms); err != nil {
		return fmt.Errorf("failed to write JSON result: %w", err)
	}

	return nil
}
