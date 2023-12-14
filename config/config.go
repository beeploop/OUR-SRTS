package config

import (
	"net"
	"strings"

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

	localAddr, err := NetDial()
	if err != nil {
		return err
	}

	var envFile map[string]string
	envFile, err = godotenv.Read()
	if err != nil {
		return err
	}

	Env = &Config{
		LocalAddr: strings.Split(localAddr, ":")[0],
		Port:      envFile["PORT"],
		DSN:       envFile["DB_DSN"],
		TempDir:   envFile["TEMP_DIR"],
		NasUrl:    envFile["NAS_URL"],
		GinMode:   envFile["GIN_MODE"],
		LogLevel:  envFile["LOG_LEVEL"],
	}

	return nil
}

func NetDial() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()

	return localAddr, nil
}
