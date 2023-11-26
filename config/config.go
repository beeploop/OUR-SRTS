package config

import "net"

type Config struct {
	Ip   string
	Port string
}

var Env *Config

const BASE_FOLDER = "documents/"

func Initialize() error {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()

	Env = &Config{
		Ip: localAddr,
        Port: ":3000",
	}

	return nil
}
