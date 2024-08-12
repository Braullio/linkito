package database

import (
	"log"
	"os"
)

type config struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	SSLMode  string
}

func loadConfig() config {
	config := config{
		Host:     getEnv("DB_HOST"),
		User:     getEnv("DB_USER"),
		Password: getEnv("DB_PASSWORD"),
		Name:     getEnv("DB_NAME"),
		Port:     getEnv("DB_PORT"),
		SSLMode:  getEnv("DB_SSLMODE"),
	}

	return config
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}
