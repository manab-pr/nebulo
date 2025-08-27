package routes

import (
	"github.com/manab-pr/nebulo/modules/auth/middleware"
	"github.com/manab-pr/nebulo/modules/search/presentation/http/handlers"
	"github.com/manab-pr/nebulo/internal/constants"

	"github.com/gin-gonic/gin"
)

func SetupSearchRoutes(router *gin.RouterGroup, handler *handlers.SearchHandler) {
	files := router.Group("/files")
	files.Use(middleware.AuthMiddleware()) // Require authentication for all search routes
	files.GET(constants.SearchFilesRoute, handler.SearchFiles)
	files.GET(constants.GetFileLocationRoute, handler.GetFileLocation)
}
