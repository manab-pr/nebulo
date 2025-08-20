package routes

import (
	"github.com/manab-pr/nebulo/modules/search/presentation/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupSearchRoutes(router *gin.RouterGroup, handler *handlers.SearchHandler) {
	files := router.Group("/files")
	files.GET("/search", handler.SearchFiles)
	files.GET("/location/:fileId", handler.GetFileLocation)
}
