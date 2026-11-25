package client

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
)

func ConnectToServer(c ClientConfig, logger *slog.Logger) error {
	serverAddress := c.ServerAddress()
	if serverAddress == "" {
		return errors.New("ino server address set")
	}
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}
	defer conn.Close()
	logger.Info("connected to server", "server", serverAddress)
	return nil
}
