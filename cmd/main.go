package main

import (
	"URLProject/configs"
	"URLProject/internal/delivery"
	"URLProject/internal/delivery/handlers"
	"URLProject/pkg/db"
)

func main() {
	config := configs.LoadConfig()
	_ = db.NewDb(config)
	authDeps := handlers.AuthServerDeps{Config: config}
	authServer := handlers.NewAuthServer(authDeps)

	delivery.StartServer(authServer)
}
