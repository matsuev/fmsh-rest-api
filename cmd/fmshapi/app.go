package main

import (
	"context"
	"fmsh-rest-api/internal/server"
	"log"
)

// ServerInterface interface
type ServerInterface interface {
	Listen() error
	Shutdown(context.Context) error
}

// Application ...
type Application struct {
	srv ServerInterface
}

// NewApplication function
func NewApplication(serverType string, routerType string) *Application {
	return &Application{
		srv: server.NewServer(serverType, routerType),
	}
}

// Start function
func (a *Application) Start() {
	if err := a.srv.Listen(); err != nil {
		log.Println(err)
	}
}

// Shutdown function
func (a *Application) Shutdown() {
	if err := a.srv.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}
}
