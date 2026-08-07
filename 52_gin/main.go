package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// Create a Gin router witd default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return Json response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start sever at port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
