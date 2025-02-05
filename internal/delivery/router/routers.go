package router

import (
	"URLProject/configs"
	"URLProject/internal/delivery/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouting(r *gin.Engine, authServer *handlers.AuthServer, linkServer *handlers.LinkServer, config *configs.Config) {
	RegisterAuthRoutes(r, authServer)
	RegisterLinkRoutes(r, linkServer, config)
}
