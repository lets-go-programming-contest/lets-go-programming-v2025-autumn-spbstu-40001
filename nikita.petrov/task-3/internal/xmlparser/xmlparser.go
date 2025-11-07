package xmlparser

import (
	"encoding/xml"
	"fmt"
	"os"
	"path"

	"github.com/Nekich06/task-3/internal/currencyrate"
	"github.com/paulrosania/go-charset/charset"
)

func WriteInfoFromInputFileToCurrRate(inputFilePath string, cbCurrencyRate *currencyrate.CurrencyRate) error {
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("can't open file %s: %w", path.Base(inputFilePath), err)
	}

	XMLDecoder := xml.NewDecoder(inputFile)
	XMLDecoder.CharsetReader = charset.NewReader

	if err := XMLDecoder.Decode(&cbCurrencyRate); err != nil {
		return fmt.Errorf("failed to decode file %s: %w", inputFile.Name(), err)
	}

	return nil
}
