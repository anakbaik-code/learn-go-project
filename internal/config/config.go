package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiKey  string
	AppPort string
	DBHost  string
	DBUser  string
	DBPass  string
	DBName  string
	LogLevel string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	return &Config{
		ApiKey:  os.Getenv("API_KEY"),
		AppPort: os.Getenv("APP_PORT"),
		DBHost:  os.Getenv("DB_HOST"),
		DBUser:  os.Getenv("DB_USER"),
		DBPass:  os.Getenv("DB_PASS"),
		DBName:  os.Getenv("DB_NAME"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}
}
