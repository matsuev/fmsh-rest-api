package server

import "context"

// ServerInterface interface
type ServerInterface interface {
	Listen() error
	Shutdown(context.Context) error
}

func NewServer(serverType string, routerType string) ServerInterface {
	switch serverType {
	case "test":
		return NewTestServer()
	default:
		return NewHttpServer(routerType)
	}
}
