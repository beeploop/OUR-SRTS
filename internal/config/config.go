package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string

	DB_USER string
	DB_PASS string
	DB_NET  string
	DB_PORT string
	DB_HOST string
	DB_NAME string

	UPLOAD_DIR string

	SECRET_KEY string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		PORT:       mustGetEnv("PORT"),
		DB_USER:    mustGetEnv("DB_USER"),
		DB_PASS:    mustGetEnv("DB_PASS"),
		DB_NET:     mustGetEnv("DB_NET"),
		DB_PORT:    mustGetEnv("DB_PORT"),
		DB_HOST:    mustGetEnv("DB_HOST"),
		DB_NAME:    mustGetEnv("DB_NAME"),
		UPLOAD_DIR: mustGetEnv("UPLOAD_DIR"),
		SECRET_KEY: mustGetEnv("SECRET_KEY"),
	}
}

func mustGetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("Environment key '%s' does not exits in environment", key))
	}

	return value
}
