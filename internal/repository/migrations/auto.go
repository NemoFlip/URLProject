package main

import (
	"URLProject/internal/entity"
	"URLProject/pkg/db"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	database := db.NewDb(os.Getenv("DSN"))
	if err := database.DB.AutoMigrate(&entity.Link{}, &entity.User{}); err != nil {
		panic(err)
	}
}
