package main

import (
	"github.com/aayushxrj/go-gin-gorm-demo/internal/config"
	"github.com/aayushxrj/go-gin-gorm-demo/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.Load()
	cfg := config.Get()

	// Initialize logger
	logger := logger.New()
	logger.Info("Starting application", map[string]interface{}{
		"name":    cfg.App.Name,
		"version": cfg.App.Version,
		"env":     cfg.App.Env,
	})

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
