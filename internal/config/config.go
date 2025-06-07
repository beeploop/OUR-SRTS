package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		PORT: mustGetEnv("PORT"),
	}
}

func mustGetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("Environment key '%s' does not exits in environment", key))
	}

	return value
}
