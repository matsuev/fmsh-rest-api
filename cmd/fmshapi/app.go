package main

import (
	"context"
	"fmsh-rest-api/internal/server"
	"log"
	"net/http"
)

// Application ...
type Application struct {
	srv *http.Server
}

// NewApplication function
func NewApplication(routerType string) *Application {
	return &Application{
		srv: server.New(routerType),
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
