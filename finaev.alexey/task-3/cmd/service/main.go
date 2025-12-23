package main

import (
	"flag"
	"os"

	"github.com/AlexeyFinaev02/task-3/internal/config"
	"github.com/AlexeyFinaev02/task-3/internal/jsonwriter"
	"github.com/AlexeyFinaev02/task-3/internal/valcurs"
	"github.com/AlexeyFinaev02/task-3/internal/xmlparser"
)

const (
	DirPerm  os.FileMode = 0o755
	FilePerm os.FileMode = 0o644
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to yaml file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err.Error())
	}

	var curs valcurs.Currency

	err = xmlparser.LoadCurrencies(cfg.InputFile, &curs)
	if err != nil {
		panic(err.Error())
	}

	curs.SortCurrenciesByValueDesc()

	err = jsonwriter.SaveJSON(cfg.OutputFile, curs.Currencies, DirPerm, FilePerm)
	if err != nil {
		panic(err.Error())
	}
}
