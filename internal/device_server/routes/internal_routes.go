package routes

import (
	"github.com/manab-pr/nebulo/internal/device_server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupInternalRoutes(router *gin.Engine, handler *handlers.InternalDeviceHandler) {
	internal := router.Group("/internal")
	internal.POST("/store", handler.StoreFile)
	internal.GET("/files/:id", handler.GetFile)
	internal.GET("/storage", handler.GetStorageInfo)
	internal.POST("/confirm/:fileId", handler.ConfirmFile)
}
