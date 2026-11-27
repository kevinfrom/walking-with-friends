package client

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"walking-with-friends/internal/shared/protocol"
)

func ConnectToServer(c ClientConfig, logger *slog.Logger) error {
	serverAddress := c.ServerAddress()
	if serverAddress == "" {
		return errors.New("no server address set")
	}
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}
	defer conn.Close()
	go func() {
		dec := json.NewDecoder(conn)
		for {
			var msg protocol.Message
			if err := dec.Decode(&msg); err != nil {
				if errors.Is(err, io.EOF) {
					logger.Info("disconnected from server")
					os.Exit(0)
				}
				logger.Error("lost connection to server", "error", err)
				os.Exit(1)
			}
			handleServerMessage(msg, logger)
		}
	}()
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		err := protocol.Write(conn, protocol.MessageTypeMove, protocol.MessagePayloadMove{})
		if err != nil {
			return fmt.Errorf("write failed: %w", err)
		}
	}
	return stdin.Err()
}

func handleServerMessage(msg protocol.Message, logger *slog.Logger) {
	switch msg.Type {
	case protocol.MessageTypeError:
		var ep protocol.ErrorPayload
		err := json.Unmarshal(msg.Payload, &ep)
		if err != nil {
			logger.Error("failed to unmarshal json from server", "error", err)
			return
		}
		logger.Error("error from server", slog.Group("error", "code", ep.Code, "message", ep.Message))
	case protocol.MessageTypeWelcome:
		var mp protocol.MessagePayloadWelcome
		err := json.Unmarshal(msg.Payload, &mp)
		if err != nil {
			logger.Error("failed to unmarshal json from server", "error", err)
			return
		}
		logger.Info("connected to server", slog.Group("player", "id", mp.PlayerID))
	}
}
