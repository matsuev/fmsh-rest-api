package router

import "net/http"

// NewRouter function
func NewRouter(routerType string) http.Handler {
	var handler http.Handler

	switch routerType {
	case "gin":
		handler = NewGinRouter()
	default:
		handler = NewHttpRouter()
	}

	return handler
}
