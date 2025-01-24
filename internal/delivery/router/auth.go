package router

import (
	"URLProject/internal/delivery/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, authServer *handlers.AuthServer) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authServer.RegisterUser)
		authGroup.POST("/login", authServer.LoginUser)
	}
}
