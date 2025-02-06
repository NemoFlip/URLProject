package router

import (
	"URLProject/configs"
	"URLProject/internal/delivery/handlers"
	"URLProject/internal/delivery/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterStatRouting(r *gin.Engine, statServer *handlers.StatServer, config *configs.Config) {
	r.GET("/stat", middleware.RequireAuthorization(config), statServer.GetStatistics)
}
