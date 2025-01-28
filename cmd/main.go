package main

import (
	"URLProject/configs"
	_ "URLProject/docs"
	"URLProject/internal/delivery"
	"URLProject/internal/delivery/handlers"
	"URLProject/internal/repository"
	"URLProject/pkg/db"
)

func main() {
	config := configs.LoadConfig()
	database := db.NewDb(config.Db.Dsn)

	// Repositories
	linkRepository := repository.NewLinkRepository(database)

	// Servers
	authDeps := handlers.AuthServerDeps{Config: config}
	authServer := handlers.NewAuthServer(authDeps)

	linkServer := handlers.NewLinkServer(linkRepository)

	delivery.StartServer(authServer, linkServer)
}
