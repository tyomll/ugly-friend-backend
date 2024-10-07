package config

import (
	_ "embed"
	"fmt"
	"log/slog"

	"gopkg.in/yaml.v3"
)

//go:embed config.yml
var embeddedConfig []byte

func MustLoad() (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal(embeddedConfig, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse embedded config: %v", err)
	}

	slog.Debug("Config init successfully")
	return &cfg, nil
}
