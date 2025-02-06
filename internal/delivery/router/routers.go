package router

import (
	"URLProject/configs"
	"URLProject/internal/delivery/handlers"
	"github.com/gin-gonic/gin"
)

// TODO: implement routing deps
func InitRouting(r *gin.Engine,
	authServer *handlers.AuthServer,
	linkServer *handlers.LinkServer,
	statServer *handlers.StatServer,
	config *configs.Config) {
	RegisterAuthRoutes(r, authServer)
	RegisterLinkRoutes(r, linkServer, config)
	RegisterStatRouting(r, statServer, config)
}
