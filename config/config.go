package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

const defaultPath = "./config/config.json"

type (
	// Config -.
	Config struct {
		RedisCluster RedisCluster `json:"redis-cluster"`
		MySQL        MySQL        `json:"mysql"`
	}

	// RedisCluster Redis-cluster -.
	RedisCluster struct {
		Addrs        []string `env-required:"true" json:"addrs"`
		Password     string   `env-required:"true" json:"password"`
		DialTimeout  int      `env-default:"10" json:"dialTimeout"`
		ReadTimeout  int      `env-default:"10" json:"readTimeout"`
		WriteTimeout int      `env-default:"10" json:"writeTimeout"`
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
