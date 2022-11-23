package app

import (
	"context"
	"fmsh-rest-api/internal/config"
	"fmsh-rest-api/internal/router"
	"fmsh-rest-api/internal/server"
	"log"
	"net/http"
)

// Application ...
type Application struct {
	srv *http.Server
}

// New function
func New(cfg *config.AppConfig) *Application {
	srv := server.New(cfg.Server)
	srv.Handler = router.New(cfg.Router)

	return &Application{
		srv: srv,
	}
}

// Start function
func (a *Application) Start() {
	if err := a.srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

// Shutdown function
func (a *Application) Shutdown() {
	if err := a.srv.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}
}
