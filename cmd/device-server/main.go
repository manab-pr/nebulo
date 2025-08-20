package main

import (
	"log"

	"github.com/manab-pr/nebulo/config"
	"github.com/manab-pr/nebulo/internal/device_server/handlers"
	"github.com/manab-pr/nebulo/internal/device_server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	logger := config.InitLogger(cfg.Server.Env)
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Printf("Failed to sync logger: %v", err)
		}
	}()

	// Set gin mode
	if cfg.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router
	router := gin.Default()

	// Initialize handlers
	deviceHandler := handlers.NewInternalDeviceHandler(cfg.Storage.Path)

	// Setup routes
	routes.SetupInternalRoutes(router, deviceHandler)

	// Start device server
	logger.Sugar().Infof("Starting device server on port %s", cfg.Device.ServerPort)
	if err := router.Run(":" + cfg.Device.ServerPort); err != nil {
		logger.Sugar().Fatalf("Device server failed to start: %v", err)
	}
}
