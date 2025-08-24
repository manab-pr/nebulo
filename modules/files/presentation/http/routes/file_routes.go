package routes

import (
	"github.com/manab-pr/nebulo/modules/auth/middleware"
	"github.com/manab-pr/nebulo/modules/files/presentation/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupFileRoutes(router *gin.RouterGroup, handler *handlers.FileHandler) {
	files := router.Group("/files")
	files.Use(middleware.AuthMiddleware()) // Require authentication for all file routes
	files.POST("/store", handler.StoreFile)
	files.GET("/:fileId", handler.GetFile)
	files.GET("", handler.GetAllFiles)
	files.DELETE("/:fileId", handler.DeleteFile)
}
