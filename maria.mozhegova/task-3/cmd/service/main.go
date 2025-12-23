package main

import (
	"flag"

	"github.com/mipxe/task-3/internal/config"
	"github.com/mipxe/task-3/internal/currency"
	"github.com/mipxe/task-3/internal/json"
	"github.com/mipxe/task-3/internal/xml"
)

const (
	DirPerm  = 0o755
	FilePerm = 0o644
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to yaml file")
	flag.Parse()

	config, err := config.ReadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	var valCurs currency.ValCurs

	err = xml.ParseXML(config.InputFile, &valCurs)
	if err != nil {
		panic(err)
	}

	valCurs.SortByValueDesc()

	err = json.WriteToJSON(valCurs.Valutes, config.OutputFile, DirPerm, FilePerm)
	if err != nil {
		panic(err)
	}
}
