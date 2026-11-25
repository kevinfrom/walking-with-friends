package main

import (
	"os"
	"walking-with-friends/internal/client"
	"walking-with-friends/internal/shared"
)

func main() {
	config := client.GetConfig()

	logger := shared.NewLogger()
	logger.Debug("config loaded")

	logger.Info("connecting to server", "server", config.ServerAddress())
	err := client.ConnectToServer(config, logger)
	if err != nil {
		logger.Error("failed to connect to server", "error", err)
		os.Exit(1)
	}
}
