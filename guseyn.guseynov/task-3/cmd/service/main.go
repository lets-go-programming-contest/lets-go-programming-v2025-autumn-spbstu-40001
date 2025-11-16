package main

import (
	"flag"
	"sort"

	"github.com/GuseynovGuseynGG/task-3/internal/config"
	"github.com/GuseynovGuseynGG/task-3/internal/currency"
	"github.com/GuseynovGuseynGG/task-3/internal/jsonwriter"
	"github.com/GuseynovGuseynGG/task-3/internal/xmlparser"
)

const (
	dirPermission  = 0o755
	filePermission = 0o644
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	var valCurs currency.ValCurs

	err = xmlparser.Parse(cfg.InputFile, &valCurs)
	if err != nil {
		panic(err)
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	jsonwriter.Write(cfg.OutputFile, valCurs.Valutes, dirPermission, filePermission)
}
