package currency

type Bank struct {
	Items []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode  int    `json:"num_code" xml:"NumCode"`
	CharCode string `json:"char_code" xml:"CharCode"`
	Value    string `xml:"Value"`
}

type outputItem struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type outputList []outputItem
