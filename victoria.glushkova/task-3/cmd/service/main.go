package main

import (
	"flag"
	"fmt"
	"github.com/vikaglushkova/task-3/internal/currency"
	"github.com/vikaglushkova/task-3/internal/json"
	"github.com/vikaglushkova/task-3/internal/xml"
)

func main() {
	inputFile := flag.String("input", "source/input_02_03_2002.xml", "Path to input XML file")
	outputFile := flag.String("output", "result/output_02_03_2002.json", "Path to output JSON file")
	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		panic("Input and output file paths are required")
	}

	valCurs, err := xml.ParseXMLFile(*inputFile)
	if err != nil {
		panic(fmt.Sprintf("Error reading XML data: %v", err))
	}

	currencies := currency.ConvertAndSort(valCurs)

	err = json.WriteToFile(*outputFile, currencies)
	if err != nil {
		panic(fmt.Sprintf("Error saving results: %v", err))
	}

	fmt.Printf("Successfully processed %d currencies. Results saved to: %s\n", len(currencies), *outputFile)
}
