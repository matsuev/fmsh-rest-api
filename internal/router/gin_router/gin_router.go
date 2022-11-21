package gin_router

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// New function
func New(logEnable bool) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())

	if logEnable {
		router.Use(gin.Logger())
	} else {
		gin.DefaultWriter = io.Discard
	}

	router.GET("/hello", handleGetHello)
	router.POST("/hello", handlePostHello)

	return router
}
