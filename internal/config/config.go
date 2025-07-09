package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		Port         string `envconfig:"SERVER_PORT" default:"8888"`
		ReadTimeout  int    `envconfig:"SERVER_READ_TIMEOUT" default:"5"`
		WriteTimeout int    `envconfig:"SERVER_WRITE_TIMEOUT" default:"10"`
		IdleTimeout  int    `envconfig:"SERVER_IDLE_TIMEOUT" default:"15"`
	}
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	var cfg Config
	if err != nil {
		cfg.Server.Port = "8888"
		cfg.Server.ReadTimeout = 5
		cfg.Server.WriteTimeout = 10
		cfg.Server.IdleTimeout = 15
		return &cfg, nil
	}

	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
