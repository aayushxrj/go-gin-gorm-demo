package main

import (
	"log"

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
	router := gin.New()

	// Setup routes

	// Start server
	addr := ":" + cfg.App.Port
	logger.Info("Server starting", map[string]interface{}{
		"port": cfg.App.Port,
		"addr": addr,
	})

	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
