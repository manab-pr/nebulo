package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/nebulo/modules/auth/middleware"
	"github.com/manab-pr/nebulo/modules/users/presentation/http/handlers"
)

func SetupUserRoutes(router *gin.RouterGroup, userHandler *handlers.UserHandler) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", userHandler.Register)
		authGroup.POST("/login", userHandler.Login)
		authGroup.POST("/verify-otp", userHandler.VerifyOTP)
	}

	userGroup := router.Group("/users")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.GET("/profile", userHandler.GetProfile)
	}
}