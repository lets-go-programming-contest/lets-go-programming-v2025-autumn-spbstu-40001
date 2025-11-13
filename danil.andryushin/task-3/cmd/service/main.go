package main

import (
	"flag"
	"sort"

	"github.com/atroxxxxxx/task-3/internal/config"
	"github.com/atroxxxxxx/task-3/internal/parsing/read"
	"github.com/atroxxxxxx/task-3/internal/parsing/write"
	"github.com/atroxxxxxx/task-3/internal/valute"
)

const DefaultPermission = 0o666

func main() {
	path := flag.String("config", "config.yaml", "config path")
	flag.Parse()

	var data config.Config

	err := read.ParseYAML(*path, &data)
	if err != nil {
		panic(err)
	}

	var valuteSlice valute.ValuteSlice

	err = read.ParseXML(data.InputFile, &valuteSlice)
	if err != nil {
		panic(err)
	}

	sort.Slice(valuteSlice.Valutes, func(i, j int) bool {
		return valuteSlice.Valutes[i].Value > valuteSlice.Valutes[j].Value
	})

	err = write.SerializeToJSON(data.OutputFile, valuteSlice.Valutes, DefaultPermission, DefaultPermission)
	if err != nil {
		panic(err)
	}
}
