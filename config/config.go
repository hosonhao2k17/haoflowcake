package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App *AppConfig
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
	}
}
