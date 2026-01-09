package main

import (
	"flag"
	"fmt"

	"github.com/A1exCRE/task-3/internal/config"
	"github.com/A1exCRE/task-3/internal/currency"
	"github.com/A1exCRE/task-3/pkg/check"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	check.Err("load config", err)

	bankData, err := currency.LoadFromFile(cfg.Input)
	check.Err("load XML", err)

	converted, err := currency.ConvertValues(bankData)
	check.Err("convert values", err)
	converted.SortDesc()

	err = converted.WriteJSONFile(cfg.Output)
	check.Err("save JSON", err)

	fmt.Println("ok")
}
