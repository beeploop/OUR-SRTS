package config

import (
	"net"

	"github.com/joho/godotenv"
)

type Config struct {
	LocalAddr  string
	Port       string
	DSN        string
	BaseFolder string
	GinMode    string
}

var Env *Config

func Initialize() error {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()

	var envFile map[string]string
	envFile, err = godotenv.Read()
	if err != nil {
		return err
	}

	Env = &Config{
		LocalAddr:  localAddr,
		Port:       envFile["PORT"],
		DSN:        envFile["DB_DSN"],
		BaseFolder: envFile["BASE_FOLDER"],
		GinMode:    envFile["GIN_MODE"],
	}

	return nil
}
