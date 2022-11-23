package http_router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// New function ...
func New() http.Handler {
	router := httprouter.New()

	router.GET("/hello", handleGetHello)
	router.POST("/hello", handlePostHello)

	return router
}
