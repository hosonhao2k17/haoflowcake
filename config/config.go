package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App *AppConfig
	Db  *DatabaseConfig
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("cannot load env")
	}
	return &Config{
		App: &AppConfig{
			AppPort: os.Getenv("APP_PORT"),
			AppEnv:  os.Getenv("APP_ENV"),
		},
		Db: &DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Port:     os.Getenv("DB_PORT"),
			Dbname:   os.Getenv("DB_NAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Sslmode:  "disable",
		},
	}
}
