package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	appPort string
	DBUrl   string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("[Error] No Env Found!")
	}

	return &Config{
		appPort: os.Getenv("APP_PORT"),
		DBUrl:   os.Getenv("DB_URL"),
	}
}
