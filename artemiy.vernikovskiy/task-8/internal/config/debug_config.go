//go:build debug

package config

import (
	_ "embed"
)

//go:embed test_debug.yaml
var configFile []byte
