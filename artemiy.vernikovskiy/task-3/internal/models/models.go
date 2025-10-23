package models

type CommaFloat float64 // please, let this live

type Settings struct {
	InputFileSetting  string `yaml:"input-file"`
	OutputFileSetting string `yaml:"output-file"`
}

type ActualData struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    CommaFloat `json:"value"     xml:"Value"`
}

type ValCurs struct {
	Valutes []ActualData `xml:"Valute"`
}
