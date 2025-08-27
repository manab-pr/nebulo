package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/manab-pr/nebulo/internal/constants"
	"github.com/manab-pr/nebulo/modules/auth/middleware"
	"github.com/manab-pr/nebulo/modules/users/presentation/http/handlers"
)

func SetupUserRoutes(router *gin.RouterGroup, userHandler *handlers.UserHandler) {
	authGroup := router.Group(constants.AuthBaseRoute)
	authGroup.POST(constants.RegisterRoute, userHandler.Register)
	authGroup.POST(constants.LoginRoute, userHandler.Login)
	authGroup.POST(constants.VerifyOTPRoute, userHandler.VerifyOTP)

	userGroup := router.Group("/users")
	userGroup.Use(middleware.AuthMiddleware())
	userGroup.GET(constants.ProfileRoute, userHandler.GetProfile)
}
