package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	router *gin.Engine
}

func NewGinRouter() *GinRouter {
	r := &GinRouter{
		router: gin.New(),
	}

	r.router.Use(gin.Logger(), gin.Recovery())

	r.router.GET("/hello", r.handleGetHello)
	r.router.POST("/hello", r.handlePostHello)

	return r
}

func (r *GinRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func (r *GinRouter) handleGetHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello!")
}

func (r *GinRouter) handlePostHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello! post request")
}
