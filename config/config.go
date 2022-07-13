package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

const defaultPath = "./config/config.json"

type (
	// Config -.
	Config struct {
		Redis Redis `json:"redis"`
		MySQL MySQL `json:"mysql"`
	}

	// Redis -.
	Redis struct {
		Name     string `env-required:"true" json:"name"`
		Port     string `env-required:"true" json:"port"`
		Password string `env-required:"true" json:"password"`
		Db       int    `env-default:"1" json:"db"`
	}
	// MySQL -.
	MySQL struct {
		Host     string `env-required:"true" json:"host"`
		Port     string `env-required:"true" json:"port"`
		User     string `env-required:"true" json:"user"`
		Password string `env-required:"true" json:"password"`
		Db       string `env-required:"true" json:"db"`
	}
)

// NewConfig returns app config.
func NewConfig(configPath string) (*Config, error) {

	if configPath == "" {
		configPath = defaultPath
	}
	cfg := &Config{}

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
