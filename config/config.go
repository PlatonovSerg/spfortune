package config

import (
	`log`
	`os`

	`github.com/joho/godotenv`
)

type Config struct {
	DSN       string
	FernetKey string
	JWTSecret string
	Port      string
}

var AppConfig Config

func LoadEnv() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	AppConfig = Config{
		DSN:       mustGetEnv("DSN"),
		FernetKey: mustGetEnv("FERNET_KEY"),
		JWTSecret: mustGetEnv("JWT_SECRET"),
		Port:      mustGetEnv("PORT"),
	}
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}
