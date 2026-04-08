package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBUrl   string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("[Error] No Env Found: ", err)
	}

	AppPort := os.Getenv("APP_PORT")
	if AppPort == "" {
		AppPort = "3000" // default port kalau kosong
	}

	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		DBUrl:   os.Getenv("DB_URL"),
	}
}
