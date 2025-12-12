//go:build dev

package config

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed dev.yaml
var configData []byte

func init() {
	loadConfig = func() (*Config, error) {
		var cfg Config
		err := yaml.Unmarshal(configData, &cfg)
		return &cfg, err
	}
}
