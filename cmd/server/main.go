package main

import (
	"walking-with-friends/internal/server"
	"walking-with-friends/internal/shared"
)

func main() {
	config := server.GetConfig()

	logger := shared.NewLogger()
	logger.Debug("config loaded")
	logger.Debug("logger ready")

	logger.Info("server starting", "port", config.Port())
}
