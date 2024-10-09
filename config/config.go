package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	LocalAddr string
	Port      string
	DSN       string
	TempDir   string
	NasUrl    string
	GinMode   string
	LogLevel  string
}

var Env *Config

func Initialize() error {
	var envFile map[string]string
	envFile, err := godotenv.Read()
	if err != nil {
		return err
	}

	Env = &Config{
		LocalAddr: envFile["LOCAL_ADDR"],
		Port:      envFile["PORT"],
		DSN:       envFile["DB_DSN"],
		TempDir:   envFile["TEMP_DIR"],
		NasUrl:    envFile["NAS_URL"],
		GinMode:   envFile["GIN_MODE"],
		LogLevel:  envFile["LOG_LEVEL"],
	}

	return nil
}
