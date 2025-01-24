package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}
type AuthConfig struct {
	SecretKey string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic("unable to load config file")
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			SecretKey: os.Getenv("JWT_SECRET_KEY"),
		},
	}
}
