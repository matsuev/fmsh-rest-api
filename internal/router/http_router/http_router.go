package http_router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// New function
func New() http.Handler {
	router := httprouter.New()

	router.GET("/hello", handleGetHello)
	router.POST("/hello", handlePostHello)

	return router
}

// handleGetHello function
func handleGetHello(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	w.Write([]byte("Hello!"))
}

// handlePostHello function
func handlePostHello(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	w.Write([]byte("Hello! post request"))
}
