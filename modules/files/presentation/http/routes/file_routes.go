package routes

import (
	"github.com/manab-pr/nebulo/modules/auth/middleware"
	"github.com/manab-pr/nebulo/modules/files/presentation/http/handlers"
	"github.com/manab-pr/nebulo/internal/constants"

	"github.com/gin-gonic/gin"
)

func SetupFileRoutes(router *gin.RouterGroup, handler *handlers.FileHandler) {
	files := router.Group(constants.FileBaseRoute)
	files.Use(middleware.AuthMiddleware()) // Require authentication for all file routes
	files.POST(constants.StoreFileRoute, handler.StoreFile)
	files.GET(constants.GetFileRoute, handler.GetFile)
	files.GET(constants.GetAllFilesRoute, handler.GetAllFiles)
	files.DELETE(constants.DeleteFileRoute, handler.DeleteFile)
}
