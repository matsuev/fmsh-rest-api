package http_router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// handleGetHello function
func handleGetHello(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	w.Write([]byte("Hello!"))
}

// handlePostHello function
func handlePostHello(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	w.Write([]byte("Hello! post request"))
}
