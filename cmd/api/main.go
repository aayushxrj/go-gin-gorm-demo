package main

import (
	"github.com/aayushxrj/go-gin-gorm-demo/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.Load()
	cfg := config.Get()

	// Initialize logger

	// Initialize database

	// Set Gin mode
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Create Gin router

	// Setup routes

	// Start server

}
