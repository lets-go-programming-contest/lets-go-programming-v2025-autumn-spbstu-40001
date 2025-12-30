
//go:build !dev

package config

import (
	_ "embed"
)

//go:embed configs/prod.yaml
var configDataProd []byte

type embedLoaderProd struct{}

func (e *embedLoaderProd) Load() ([]byte, error) {
	return configDataProd, nil
}

func NewLoader() Loader {
	return &embedLoaderProd{}
}
