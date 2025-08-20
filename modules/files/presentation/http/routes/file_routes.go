package routes

import (
	"github.com/manab-pr/nebulo/modules/files/presentation/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupFileRoutes(router *gin.RouterGroup, handler *handlers.FileHandler) {
	files := router.Group("/files")
	files.POST("/store", handler.StoreFile)
	files.GET("/:fileId", handler.GetFile)
	files.GET("", handler.GetAllFiles)
	files.DELETE("/:fileId", handler.DeleteFile)
}
