package routes

import (
	"github.com/manab-pr/nebulo/modules/devices/presentation/http/handlers"

	"github.com/gin-gonic/gin"
)

func SetupDeviceRoutes(router *gin.RouterGroup, handler *handlers.DeviceHandler) {
	devices := router.Group("/devices")
	devices.POST("/register", handler.RegisterDevice)
	devices.POST("/heartbeat", handler.Heartbeat)
	devices.GET("", handler.GetDevices)
	devices.DELETE("/:id", handler.DeleteDevice)
}
