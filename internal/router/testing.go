package router

import "net/http"

// setupRouter function
func setupRouter(mode string) http.Handler {
	cfg := &Config{
		Mode: mode,
	}

	return New(cfg)
}
