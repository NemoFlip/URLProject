package delivery

import (
	"URLProject/internal/delivery/handlers"
	"URLProject/internal/delivery/router"
	"github.com/gin-gonic/gin"
)

func StartServer(authServer *handlers.AuthServer, linkServer *handlers.LinkServer) {
	r := gin.Default()
	router.InitRouting(r, authServer, linkServer)

	if err := r.Run(":8080"); err != nil {
		panic("unable to run server on port 8080")
	}
}
