package router

import (
	"fmsh-rest-api/internal/router/gin_router"
	"fmsh-rest-api/internal/router/http_router"
	"net/http"
)

const (
	MODE_GIN = "gin"
)

// Config ...
type Config struct {
	Mode string `yaml:"mode"`
}

// New function
func New(cfg *Config) http.Handler {
	switch cfg.Mode {
	case MODE_GIN:
		return gin_router.New()
	default:
		return http_router.New()
	}
}
