//go:build dev

package config

import (
	_ "embed"
	"log"

	"gopkg.in/yaml.v3"
)

//go:embed dev.yaml
var devConfigFile []byte

func init() {
	err := yaml.Unmarshal(devConfigFile, &Current)
	if err != nil {
		log.Fatalf("Error parsing dev config: %v", err)
	}
}
