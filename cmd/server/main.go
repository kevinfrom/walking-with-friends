package main

import (
	"os"
	"walking-with-friends/internal/server"
	"walking-with-friends/internal/shared"
)

func main() {
	config := server.GetConfig()

	logger := shared.NewLogger()
	logger.Debug("config loaded")

	logger.Debug("starting server", "port", config.Port())
	err := server.StartServer(config, logger)
	if err != nil {
		logger.Error("failed to start server", "error", err.Error())
		os.Exit(1)
	}
}
