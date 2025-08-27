package routes

import (
	"github.com/manab-pr/nebulo/modules/auth/middleware"
	"github.com/manab-pr/nebulo/modules/storage/presentation/http/handlers"
	"github.com/manab-pr/nebulo/internal/constants"

	"github.com/gin-gonic/gin"
)

func SetupStorageRoutes(router *gin.RouterGroup, handler *handlers.StorageHandler) {
	storage := router.Group(constants.StorageBaseRoute)
	storage.Use(middleware.AuthMiddleware()) // Require authentication for all storage routes
	storage.GET(constants.GetStorageSummaryRoute, handler.GetStorageSummary)
	storage.GET(constants.GetDeviceStorageRoute, handler.GetDeviceStorage)
}
