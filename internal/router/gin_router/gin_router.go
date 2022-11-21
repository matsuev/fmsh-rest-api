package gin_router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// New function
func New() http.Handler {

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/hello", handleGetHello)
	router.POST("/hello", handlePostHello)

	return router
}

func handleGetHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello!")
}

func handlePostHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello! post request")
}
