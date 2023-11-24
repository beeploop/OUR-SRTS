package config

import "net"

type Config struct {
	Ip string
}

var Env *Config

func Initialize() error {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()

	Env = &Config{
		Ip: localAddr,
	}

	return nil
}
