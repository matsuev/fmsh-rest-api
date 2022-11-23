package gin_router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleGetHello function
func handleGetHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello!")
}

// handlePostHello function
func handlePostHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello! post request")
}
