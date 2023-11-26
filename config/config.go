package config

import (
	"net"

	"github.com/joho/godotenv"
)

type Config struct {
	Ip         string
	Port       string
	DSN        string
	BaseFolder string
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
		Ip:         localAddr,
		Port:       envFile["PORT"],
		DSN:        envFile["DB_DSN"],
		BaseFolder: envFile["BASE_FOLDER"],
	}

	return nil
}
