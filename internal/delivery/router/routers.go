package router

import (
	"URLProject/internal/delivery/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouting(r *gin.Engine, authServer *handlers.AuthServer, linkServer *handlers.LinkServer) {
	RegisterAuthRoutes(r, authServer)
	RegisterLinkRoutes(r, linkServer)
}
