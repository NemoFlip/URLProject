package main

import (
	"URLProject/internal/entity"
	"URLProject/internal/stat"
	"URLProject/pkg/db"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	database := db.NewDb(os.Getenv("DSN"))
	if err := database.DB.AutoMigrate(&entity.Link{}, &entity.User{}, &stat.Stat{}); err != nil {
		panic(err)
	}
}
