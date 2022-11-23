package config

import (
	"fmsh-rest-api/internal/router"
	"fmsh-rest-api/internal/server"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	DEFAULT_CONFIG_PATH = "./configs/config.yml"
)

// AppConfig ...
type AppConfig struct {
	Server *server.Config `yaml:"server"`
	Router *router.Config `yaml:"router"`
}

// New function
func New(path ...string) (*AppConfig, error) {
	cfgPath := DEFAULT_CONFIG_PATH
	if len(path) > 0 {
		cfgPath = path[0]
	}

	cfg := &AppConfig{}

	if err := cleanenv.ReadConfig(cfgPath, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
