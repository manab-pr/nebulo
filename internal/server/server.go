package server

import (
	"github.com/manab-pr/nebulo/config"
	"github.com/manab-pr/nebulo/container"
	deviceRoutes "github.com/manab-pr/nebulo/modules/devices/presentation/http/routes"
	fileRoutes "github.com/manab-pr/nebulo/modules/files/presentation/http/routes"
	searchRoutes "github.com/manab-pr/nebulo/modules/search/presentation/http/routes"
	storageRoutes "github.com/manab-pr/nebulo/modules/storage/presentation/http/routes"
	transferRoutes "github.com/manab-pr/nebulo/modules/transfers/presentation/http/routes"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	router    *gin.Engine
	config    *config.Config
	logger    *zap.Logger
	container *container.AppContainer
}

func NewServer(cfg *config.Config, logger *zap.Logger, appContainer *container.AppContainer) *Server {
	return &Server{
		config:    cfg,
		logger:    logger,
		container: appContainer,
	}
}

func (s *Server) SetupRoutes() {
	// Set gin mode
	if s.config.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router
	s.router = gin.Default()

	// Add basic middleware
	s.router.Use(gin.Logger())
	s.router.Use(gin.Recovery())

	// Health check endpoint
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": s.config.Server.Name,
		})
	})

	// API v1 routes
	v1 := s.router.Group("/api/v1")
	// Setup module routes
	deviceRoutes.SetupDeviceRoutes(v1, s.container.DeviceHandler)
	fileRoutes.SetupFileRoutes(v1, s.container.FileHandler)
	transferRoutes.SetupTransferRoutes(v1, s.container.TransferHandler)
	storageRoutes.SetupStorageRoutes(v1, s.container.StorageHandler)
	searchRoutes.SetupSearchRoutes(v1, s.container.SearchHandler)
}

func (s *Server) Run() error {
	s.logger.Sugar().Infof("Starting server on port %s", s.config.Server.Port)
	return s.router.Run(":" + s.config.Server.Port)
}
