package main

import (
	"URLProject/configs"
	_ "URLProject/docs"
	"URLProject/internal/delivery"
	"URLProject/internal/delivery/handlers"
	"URLProject/internal/delivery/middleware"
	"URLProject/internal/delivery/router"
	"URLProject/internal/delivery/services"
	"URLProject/internal/repository"
	"URLProject/pkg/db"
	"URLProject/pkg/event"
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	config := configs.LoadConfig()
	database := db.NewDb(config.Db.Dsn)
	eventBus := event.NewEventBus()

	// Repositories
	linkRepository := repository.NewLinkRepository(database)
	userRepository := repository.NewUserRepository(database)
	statRepository := repository.NewStatRepository(database)

	// Services
	authService := services.NewAuthService(userRepository)
	statService := services.NewStatService(&services.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	go statService.AddClick()

	// Servers
	authServer := handlers.NewAuthServer(handlers.AuthServerDeps{Config: config, AuthService: authService})
	linkServer := handlers.NewLinkServer(linkRepository, eventBus)
	statServer := handlers.NewStatServer(statRepository)

	r := gin.Default()
	r.Use(middleware.CORS())

	router.InitRouting(r, authServer, linkServer, statServer, config)

	return r
}

func main() {
	config := configs.LoadConfig()
	database := db.NewDb(config.Db.Dsn)
	eventBus := event.NewEventBus()

	// Repositories
	linkRepository := repository.NewLinkRepository(database)
	userRepository := repository.NewUserRepository(database)
	statRepository := repository.NewStatRepository(database)

	// Services
	authService := services.NewAuthService(userRepository)
	statService := services.NewStatService(&services.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	// Servers
	authServer := handlers.NewAuthServer(handlers.AuthServerDeps{Config: config, AuthService: authService})
	linkServer := handlers.NewLinkServer(linkRepository, eventBus)
	statServer := handlers.NewStatServer(statRepository)

	go statService.AddClick()

	delivery.StartServer(authServer, linkServer, statServer, config)
}
