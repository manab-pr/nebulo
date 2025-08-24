package routes

import (
	"github.com/manab-pr/nebulo/modules/auth/middleware"
	"github.com/manab-pr/nebulo/modules/storage/presentation/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupStorageRoutes(router *gin.RouterGroup, handler *handlers.StorageHandler) {
	storage := router.Group("/storage")
	storage.Use(middleware.AuthMiddleware()) // Require authentication for all storage routes
	{
		storage.GET("/summary", handler.GetStorageSummary)
		storage.GET("/device/:deviceId", handler.GetDeviceStorage)
	}
}
