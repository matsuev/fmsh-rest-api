package server

import (
	"fmsh-rest-api/internal/router"
	"net/http"
)

// New function ...
func New(routerType string) *http.Server {
	return &http.Server{
		Addr:    ":8081",
		Handler: router.New(routerType),
	}
}
