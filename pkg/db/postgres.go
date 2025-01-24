package db

import (
	"URLProject/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Db struct {
	DB *gorm.DB
}

func NewDb(config *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(config.Db.Dsn), &gorm.Config{})
	if err != nil {
		log.Printf("unable to connect to database: %s", err)
		return nil
	}
	return &Db{DB: db}
}
