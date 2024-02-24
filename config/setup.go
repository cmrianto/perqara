package config

import (
	"os"

	"gopkg.in/gcfg.v1"
)

func Setup() (*Config, error) {
	cfgFile := "config.toml"
	if os.Getenv("ENV") == "local" {
		cfgFile = "config.toml"
	}
	cfg := &Config{}

	err := gcfg.ReadFileInto(cfg, cfgFile)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
