package main

import (
	"flag"
	"sort"

	"github.com/LeeLisssa/task-3/internal/config"
	"github.com/LeeLisssa/task-3/internal/jsonwriter"
	"github.com/LeeLisssa/task-3/internal/types"
	"github.com/LeeLisssa/task-3/internal/xmlparser"
)

const (
	dirPermission  = 0o755
	filePermission = 0o600
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to configuration file")
	flag.Parse()

	cfg, err := config.ParseYaml(*configPath)
	if err != nil {
		panic(err)
	}

	var typesList types.Rates

	err = xmlparser.ParseXML(cfg.InputFilePath, &typesList)
	if err != nil {
		panic(err)
	}

	sort.Slice(typesList.Data, func(i, j int) bool {
		return typesList.Data[i].Value > typesList.Data[j].Value
	})

	err = jsonwriter.ParseJSON(cfg.OutputFilePath, typesList.Data, dirPermission, filePermission)
	if err != nil {
		panic(err)
	}
}
