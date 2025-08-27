package routes

import (
	"github.com/manab-pr/nebulo/modules/auth/middleware"
	"github.com/manab-pr/nebulo/modules/devices/presentation/http/handlers"
	"github.com/manab-pr/nebulo/internal/constants"

	"github.com/gin-gonic/gin"
)

func SetupDeviceRoutes(router *gin.RouterGroup, handler *handlers.DeviceHandler) {
	devices := router.Group(constants.DeviceBaseRoute)
	devices.Use(middleware.AuthMiddleware()) // Require authentication for all device routes
	devices.POST(constants.RegisterDeviceRoute, handler.RegisterDevice)
	devices.POST(constants.HeartbeatRoute, handler.Heartbeat)
	devices.GET(constants.GetDevicesRoute, handler.GetDevices)
	devices.DELETE(constants.DeleteDeviceRoute, handler.DeleteDevice)
}
