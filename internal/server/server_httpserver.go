package server

import (
	"context"
	"fmsh-rest-api/internal/router"
	"log"
	"net/http"
)

// HttpServer ...
type HttpServer struct {
	srv *http.Server
}

// NewHttpServer function
func NewHttpServer(routerType string) *HttpServer {
	return &HttpServer{
		srv: &http.Server{
			Addr:    "localhost:8081",
			Handler: router.NewRouter(routerType),
		},
	}
}

// Listen function
func (s *HttpServer) Listen() error {
	log.Println("server listening on localhost:8081")
	return s.srv.ListenAndServe()
}

// Shutdown function
func (s *HttpServer) Shutdown(ctx context.Context) error {
	log.Println("server shutdown")
	return s.srv.Shutdown(ctx)
}
