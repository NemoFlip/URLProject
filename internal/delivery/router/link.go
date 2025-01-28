package router

import (
	"URLProject/internal/delivery/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterLinkRoutes(r *gin.Engine, linkServer *handlers.LinkServer) {
	linkGroup := r.Group("/link")
	{
		linkGroup.POST("", linkServer.Create)
		linkGroup.PATCH("/:id", linkServer.Update)
		linkGroup.DELETE("", linkServer.Delete)
		linkGroup.GET("/:hash", linkServer.GoTo)
	}
}
