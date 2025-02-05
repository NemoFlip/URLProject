package main

import (
	"URLProject/configs"
	_ "URLProject/docs"
	"URLProject/internal/delivery"
	"URLProject/internal/delivery/handlers"
	"URLProject/internal/delivery/services"
	"URLProject/internal/repository"
	"URLProject/pkg/db"
)

func main() {
	config := configs.LoadConfig()
	database := db.NewDb(config.Db.Dsn)

	// Repositories
	linkRepository := repository.NewLinkRepository(database)
	userRepository := repository.NewUserRepository(database)

	// Services
	authService := services.NewAuthService(userRepository)

	// Servers
	authDeps := handlers.AuthServerDeps{Config: config, AuthService: authService}
	authServer := handlers.NewAuthServer(authDeps)

	linkServer := handlers.NewLinkServer(linkRepository)

	delivery.StartServer(authServer, linkServer, config)
}
