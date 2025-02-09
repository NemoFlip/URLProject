package router

import (
	"URLProject/configs"
	"URLProject/internal/delivery/handlers"
	"URLProject/internal/delivery/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterLinkRoutes(r *gin.Engine, linkServer *handlers.LinkServer, config *configs.Config) {
	linkGroup := r.Group("/link")
	{
		linkGroup.POST("", middleware.RequireAuthorization(config), linkServer.Create)
		linkGroup.PATCH("/:id", middleware.RequireAuthorization(config), linkServer.Update)
		linkGroup.DELETE("/:id", middleware.RequireAuthorization(config), linkServer.Delete)
		linkGroup.GET("/:hash", linkServer.GoTo)
		linkGroup.GET("", middleware.RequireAuthorization(config), linkServer.GetAll)
	}
}
