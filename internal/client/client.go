package client

import (
	"bufio"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
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
	server := bufio.NewReader(conn)
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		line := stdin.Text()
		if _, err := fmt.Fprintf(conn, "%s\n", line); err != nil {
			return fmt.Errorf("write failed: %w", err)
		}
		reply, err := server.ReadString('\n')
		if err != nil {
			return fmt.Errorf("read failed: %w", err)
		}
		logger.Debug("received reply", "reply", reply)
	}
	return stdin.Err()
}
