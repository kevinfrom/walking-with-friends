package server

import (
	"os"
	"strconv"
	"walking-with-friends/internal/shared"
)

const (
	DEFAULT_MAX_PLAYERS int = 4
)

type ServerConfig struct {
	port       string
	maxPlayers int
	baseConfig shared.BaseConfig
}

func GetConfig() ServerConfig {
	port := os.Getenv("PORT")
	maxPlayers, err := strconv.Atoi(os.Getenv("MAX_PLAYERS"))
	if err != nil || maxPlayers <= 0 {
		maxPlayers = DEFAULT_MAX_PLAYERS
	}
	return ServerConfig{
		port:       port,
		maxPlayers: maxPlayers,
		baseConfig: shared.NewBaseConfig(),
	}
}

func (c ServerConfig) Port() string {
	return c.port
}

func (c ServerConfig) MaxPlayers() int {
	return c.maxPlayers
}

func (c ServerConfig) BaseConfig() shared.BaseConfig {
	return c.baseConfig
}
