package iocurrency

type ValCurs struct {
	Valutes []struct {
		NumCode  int       `json:"num_code"  xml:"NumCode"`
		CharCode string    `json:"char_code" xml:"CharCode"`
		ValueStr UserFloat `json:"value"     xml:"Value"`
	} `xml:"Valute"`
}
