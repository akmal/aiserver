package server

import (
	"github.com/akmal/aiserver/handlers"
	"github.com/gin-gonic/gin"
)

/**
 * StartServer initializes and starts the Gin server.
 */
func StartServer(debug bool) {
	r := gin.Default()

	r.POST("/prompt", func(c *gin.Context) {
		handlers.PromptHandler(c, c.Query("debug") == "true")
	})

	r.Run(":8080")
}
