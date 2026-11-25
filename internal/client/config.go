package client

import "os"

type ClientConfig struct {
	serverAddress string
}

func GetConfig() ClientConfig {
	return ClientConfig{
		serverAddress: os.Getenv("SERVER_ADDRESS"),
	}
}

func (c ClientConfig) ServerAddress() string {
	return c.serverAddress
}
