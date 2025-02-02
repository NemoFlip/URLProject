package router

import (
	"URLProject/internal/delivery/handlers"
	"URLProject/internal/delivery/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterLinkRoutes(r *gin.Engine, linkServer *handlers.LinkServer) {
	linkGroup := r.Group("/link")
	{
		linkGroup.POST("", linkServer.Create)
		linkGroup.PATCH("/:id", middleware.RequireAuthorization(), linkServer.Update)
		linkGroup.DELETE("/:id", linkServer.Delete)
		linkGroup.GET("/:hash", linkServer.GoTo)
	}
}
