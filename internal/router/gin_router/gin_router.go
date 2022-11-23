package gin_router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// New function ...
func New() http.Handler {

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/hello", handleGetHello)
	router.POST("/hello", handlePostHello)

	return router
}
