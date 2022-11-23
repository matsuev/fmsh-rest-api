package router

import (
	"fmsh-rest-api/internal/router/gin_router"
	"fmsh-rest-api/internal/router/http_router"
	"net/http"
)

// New function
func New(routerType string) http.Handler {
	switch routerType {
	case "gin":
		return gin_router.New()
	default:
		return http_router.New()
	}
}
