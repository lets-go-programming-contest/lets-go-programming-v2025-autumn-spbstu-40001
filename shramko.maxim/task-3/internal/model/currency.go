package model

type Currency struct {
	NumericCode int     `json:"num_code"  xml:"NumCode"`
	CharCode    string  `json:"char_code" xml:"CharCode"`
	Value       float64 `json:"value"     xml:"Value"`
}

type CurrencyCollection struct {
	CurrencyItems []Currency `xml:"Valute"`
}
