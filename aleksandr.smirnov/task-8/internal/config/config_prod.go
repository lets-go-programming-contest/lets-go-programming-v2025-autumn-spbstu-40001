//go:build !dev

package config

import _ "embed"

//go:embed prod.yaml
var prodData []byte

func Init() (*AppConfig, error) {
	return ParseConfig(prodData)
}
