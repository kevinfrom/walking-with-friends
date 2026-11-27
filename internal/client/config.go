package client

import (
	"os"
	"walking-with-friends/internal/shared"
)

type ClientConfig struct {
	serverAddress string
	baseConfig    shared.BaseConfig
}

func GetConfig() ClientConfig {
	return ClientConfig{
		serverAddress: os.Getenv("SERVER_ADDRESS"),
		baseConfig:    shared.NewBaseConfig(),
	}
}

func (c ClientConfig) ServerAddress() string {
	return c.serverAddress
}

func (c ClientConfig) BaseConfig() shared.BaseConfig {
	return c.baseConfig
}
