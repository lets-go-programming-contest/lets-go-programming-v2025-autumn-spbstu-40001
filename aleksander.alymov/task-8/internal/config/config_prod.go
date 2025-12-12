//go:build !dev || prod

package config

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed prod.yaml
var configData []byte

func init() {
	loadConfig = func() (*Config, error) {
		var cfg Config
		err := yaml.Unmarshal(configData, &cfg)
		return &cfg, err
	}
}
