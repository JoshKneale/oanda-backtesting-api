package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not defined", key)
	}
	return value
}
