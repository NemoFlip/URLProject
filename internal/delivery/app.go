package delivery

import (
	"URLProject/configs"
	"URLProject/internal/delivery/handlers"
	"URLProject/internal/delivery/middleware"
	"URLProject/internal/delivery/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title URLProject
// @description Project for shorting the URLS
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /
func StartServer(authServer *handlers.AuthServer, linkServer *handlers.LinkServer, statServer *handlers.StatServer, config *configs.Config) {
	r := gin.Default()

	r.Use(middleware.CORS())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.InitRouting(r, authServer, linkServer, statServer, config)

	if err := r.Run(":8080"); err != nil {
		panic("unable to run server on port 8080")
	}
}
