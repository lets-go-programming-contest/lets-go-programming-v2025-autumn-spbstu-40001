package read

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func ParseYAML(path string, obj any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("cannot read config file: %w", err)
	}

	err = yaml.Unmarshal(data, obj)
	if err != nil {
		return fmt.Errorf("yaml unarshallig failed: %w", err)
	}

	return nil
}
