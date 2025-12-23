package xmlparser

import (
	"encoding/xml"
	"fmt"
	"os"
	"path"

	"github.com/paulrosania/go-charset/charset"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func WriteInfoFromInputFileToCurrRate[T any](inputFilePath string, cbCurrencyRate *T) error {
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("can't open file %s: %w", path.Base(inputFilePath), err)
	}

	defer func() {
		err := inputFile.Close()
		panicIfErr(err)
	}()

	XMLDecoder := xml.NewDecoder(inputFile)
	XMLDecoder.CharsetReader = charset.NewReader

	if err := XMLDecoder.Decode(&cbCurrencyRate); err != nil {
		return fmt.Errorf("failed to decode file %s: %w", inputFile.Name(), err)
	}

	return nil
}
