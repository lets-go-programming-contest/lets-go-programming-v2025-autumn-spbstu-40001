//go:build !dev

package config

import (
	_ "embed"
	"log"

	"gopkg.in/yaml.v3"
)

//go:embed prod.yaml
var prodConfigFile []byte

func init() {
	err := yaml.Unmarshal(prodConfigFile, &Current)
	if err != nil {
		log.Fatalf("Error parsing prod config: %v", err)
	}
}
