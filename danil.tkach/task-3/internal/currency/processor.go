package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

const (
	DirPerms  os.FileMode = 0o755
	FilePerms os.FileMode = 0o644
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}
type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Value    string `xml:"Value"`
}

type Currency struct {
	NumCode  int
	CharCode string
	Value    float64
}
type ByValue []Currency

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].Value > a[j].Value }

type CurrencyOutput struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func readValCurs(inputFile string) (*ValCurs, error) {
	xmlData, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read XML-file %s: %w", inputFile, err)
	}

	dataRdr := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(dataRdr)

	decoder.CharsetReader = func(c string, input io.Reader) (io.Reader, error) {
		return charset.NewReader(input, c)
	}

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("failed to parse XML-file: %w", err)
	}

	return &valCurs, nil
}

func transformAndSort(valCurs *ValCurs) ([]Currency, error) {
	currencies := make([]Currency, 0, len(valCurs.Valutes))

	for _, valute := range valCurs.Valutes {
		valueStr := strings.Replace(valute.Value, ",", ".", 1)

		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("failed convert '%s' to number: %w", valute.Value, err)
		}

		realValue := value / float64(valute.Nominal)

		currencies = append(currencies, Currency{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    realValue,
		})
	}

	sort.Sort(ByValue(currencies))

	return currencies, nil
}

func writeResult(currencies []Currency, outputFile string) error {
	outputDir := filepath.Dir(outputFile)
	if err := os.MkdirAll(outputDir, DirPerms); err != nil {
		return fmt.Errorf("failed to create a dir %s: %w", outputDir, err)
	}

	outputData := make([]CurrencyOutput, 0, len(currencies))
	for _, c := range currencies {
		outputData = append(outputData, CurrencyOutput(c))
	}

	jsonData, err := json.MarshalIndent(outputData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to read to json: %w", err)
	}

	if err := os.WriteFile(outputFile, jsonData, FilePerms); err != nil {
		return fmt.Errorf("failed write file %s: %w", outputFile, err)
	}

	return nil
}

func Process(inputFile, outputFile string) error {
	valCurs, err := readValCurs(inputFile)
	if err != nil {
		return err
	}

	currencies, err := transformAndSort(valCurs)
	if err != nil {
		return err
	}

	if err := writeResult(currencies, outputFile); err != nil {
		return err
	}

	return nil
}
