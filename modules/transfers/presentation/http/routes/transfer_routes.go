package routes

import (
	"github.com/manab-pr/nebulo/modules/auth/middleware"
	"github.com/manab-pr/nebulo/modules/transfers/presentation/http/handlers"
	"github.com/manab-pr/nebulo/internal/constants"

	"github.com/gin-gonic/gin"
)

func SetupTransferRoutes(router *gin.RouterGroup, handler *handlers.TransferHandler) {
	transfers := router.Group(constants.TransferBaseRoute)
	transfers.Use(middleware.AuthMiddleware()) // Require authentication for all transfer routes
	transfers.GET(constants.GetPendingTransfersRoute, handler.GetPendingTransfers)
	transfers.POST(constants.CompleteTransferRoute, handler.CompleteTransfer)
	transfers.DELETE(constants.CancelTransferRoute, handler.CancelTransfer)
}
