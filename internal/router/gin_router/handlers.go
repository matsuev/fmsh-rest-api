package gin_router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello!")
}

func handlePostHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello! post request")
}
