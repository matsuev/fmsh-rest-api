package router

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HttpRouter ...
type HttpRouter struct {
	router *httprouter.Router
}

func NewHttpRouter() *HttpRouter {
	r := &HttpRouter{
		router: httprouter.New(),
	}

	r.router.GET("/hello", r.handleGetHello)
	r.router.POST("/hello", r.handlePostHello)

	return r
}

func (r *HttpRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func (r *HttpRouter) handleGetHello(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	io.WriteString(w, "Hello!")
}

func (r *HttpRouter) handlePostHello(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	io.WriteString(w, "Hello! post request")
}
