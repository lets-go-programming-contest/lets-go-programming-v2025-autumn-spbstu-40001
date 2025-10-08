package currency;

import "os";
import "fmt";
import "strings";
import "strconv";
import "encoding/xml";
import "golang.org/x/net/html/charset";

type Currency struct {
	XMLName xml.Name `xml:"Valute"`;
	NumCode uint `json:"num_code"`;
	CharCode string `json:"char_code"`;
	ValueStr string `xml:"Value" json:"-"`;
	Value float32 `xml:"-" json:"value"`;
}
type CurrencyRates struct {
	XMLName xml.Name `xml:"ValCurs"`;
	Rates []Currency `xml:"Valute"`;
}

func Prepare(rates *CurrencyRates) error {
	for idx := range(len(rates.Rates)) {
		value, err := strconv.ParseFloat(strings.ReplaceAll(rates.Rates[idx].ValueStr, ",", "."), 32);
		if (err != nil) {
			return fmt.Errorf("failed to parse rate value: %w", err);
		}
		rates.Rates[idx].Value = float32(value);
	}
	return nil;
}

func ParseXml(xmlPath string) (CurrencyRates, error) {
	var result CurrencyRates;

	xmlFile, err := os.Open(xmlPath);
	if (err != nil) {
		return result, fmt.Errorf("failed to open currency list xml file: %w", err);
	}
	defer xmlFile.Close();

	decoder := xml.NewDecoder(xmlFile);
	decoder.CharsetReader = charset.NewReaderLabel;

	err = decoder.Decode(&result);
	if (err != nil) {
		return result, fmt.Errorf("failed to parse currency list xml file: %w", err);
	}

	return result, nil;
}
