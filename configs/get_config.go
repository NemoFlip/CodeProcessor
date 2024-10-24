package configs

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func GetConfig() (*Config, error) {
	var cfg Config
	f, err := os.Open("./configs/configs.yml")
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode congig: %s", err)
	}
	return &cfg, nil
}
