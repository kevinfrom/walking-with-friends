package server

import (
	"os"
)

type ServerConfig struct {
	port string
}

func GetConfig() ServerConfig {
	return ServerConfig{
		port: os.Getenv("PORT"),
	}
}

func (c ServerConfig) Port() string {
	return c.port
}
