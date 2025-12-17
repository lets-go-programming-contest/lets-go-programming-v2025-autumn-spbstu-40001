//go:build prod

package config

import (
	_ "embed"
)

//go:embed test_prod.yaml
var configFile []byte
