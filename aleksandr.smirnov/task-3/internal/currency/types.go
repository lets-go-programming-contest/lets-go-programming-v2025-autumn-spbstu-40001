package currency

type Bank struct {
	Items []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode  int    `xml:"NumCode" json:"num_code"`
	CharCode string `xml:"CharCode" json:"char_code"`
	Value    string `xml:"Value"`
}

type outputItem struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type outputList []outputItem
