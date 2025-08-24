package routes

import (
	"github.com/manab-pr/nebulo/modules/auth/middleware"
	"github.com/manab-pr/nebulo/modules/devices/presentation/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupDeviceRoutes(router *gin.RouterGroup, handler *handlers.DeviceHandler) {
	devices := router.Group("/devices")
	devices.Use(middleware.AuthMiddleware()) // Require authentication for all device routes
	devices.POST("/register", handler.RegisterDevice)
	devices.POST("/heartbeat", handler.Heartbeat)
	devices.GET("", handler.GetDevices)
	devices.DELETE("/:id", handler.DeleteDevice)
}
