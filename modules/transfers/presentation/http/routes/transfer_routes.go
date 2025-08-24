package routes

import (
	"github.com/manab-pr/nebulo/modules/auth/middleware"
	"github.com/manab-pr/nebulo/modules/transfers/presentation/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupTransferRoutes(router *gin.RouterGroup, handler *handlers.TransferHandler) {
	transfers := router.Group("/transfers")
	transfers.Use(middleware.AuthMiddleware()) // Require authentication for all transfer routes
	transfers.GET("/pending/:deviceId", handler.GetPendingTransfers)
	transfers.POST("/complete", handler.CompleteTransfer)
	transfers.DELETE("/:id", handler.CancelTransfer)
}
