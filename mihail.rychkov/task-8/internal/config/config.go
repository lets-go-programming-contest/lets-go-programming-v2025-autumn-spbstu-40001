package config;

import _ "gopkg.in/yaml.v3";

type Config struct {
	Env    string `yaml:"environment"`
	LogLvl string `yaml:"log_level"`
}
func GetActive() (Config, error) {
	return Config{"test", "test-lvl"}, nil;
}
