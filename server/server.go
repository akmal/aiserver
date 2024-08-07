package server

import (
	"fmt"
	"github.com/akmal/aiserver/handlers"
	"github.com/gin-gonic/gin"
)

/**
 * StartServer initializes and starts the Gin server.
 */
func StartServer(debug bool) {
	if debug {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()

	r.POST("/prompt", func(c *gin.Context) {
		handlers.PromptHandler(c, c.Query("debug") == "true")
	})

	fmt.Println("AI Server started listening on port :8080")

	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %s", err)
		return
	}
}
