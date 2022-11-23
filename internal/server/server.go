package server

import (
	"fmt"
	"net/http"
)

// Config ...
type Config struct {
	Addr string `yaml:"addr"`
	Port string `yaml:"port"`
}

// New function
func New(cfg *Config) *http.Server {
	return &http.Server{
		Addr: fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port),
	}
}
