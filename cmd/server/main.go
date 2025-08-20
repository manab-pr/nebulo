package main

import (
	"log"

	"github.com/manab-pr/nebulo/config"
	"github.com/manab-pr/nebulo/container"
	"github.com/manab-pr/nebulo/internal/server"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	logger := config.InitLogger(cfg.Server.Env)
	defer logger.Sync()

	// Connect to MongoDB
	db, err := config.ConnectMongoDB(cfg)
	if err != nil {
		logger.Sugar().Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Connect to Redis (optional, can be nil for now)
	redis := config.ConnectRedis(cfg)
	if redis == nil {
		logger.Sugar().Warn("Redis connection failed, continuing without Redis")
	}

	// Initialize app container
	appContainer := container.NewAppContainer(db, redis, cfg, logger)

	// Initialize server
	srv := server.NewServer(cfg, logger, appContainer)
	srv.SetupRoutes()

	// Start server
	log.Fatal(srv.Run())
}
