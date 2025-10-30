package main

import (
	"flag"
	"sort"

	"github.com/atroxxxxxx/task-3/internal/parsing/json"
	"github.com/atroxxxxxx/task-3/internal/parsing/xml"
	"github.com/atroxxxxxx/task-3/internal/parsing/yaml"
	"github.com/atroxxxxxx/task-3/internal/valute"
)

const Permission = 0o666

func main() {
	path := flag.String("config", "config.yaml", "config path")
	flag.Parse()

	data, err := yaml.Unmarshall(*path)
	if err != nil {
		panic(err)
	}

	var valuteSlice valute.ValuteSlice

	err = xml.Unmarshall(data.InputFile, &valuteSlice)
	if err != nil {
		panic(err)
	}

	sort.Slice(valuteSlice.Valutes, func(i, j int) bool {
		return valuteSlice.Valutes[i].Value > valuteSlice.Valutes[j].Value
	})

	err = json.WriteToFile(data.OutputFile, valuteSlice.Valutes, Permission)
	if err != nil {
		panic(err)
	}
}
