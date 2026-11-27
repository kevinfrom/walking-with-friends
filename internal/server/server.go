package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"walking-with-friends/internal/shared/protocol"
)

func StartServer(config ServerConfig, logger *slog.Logger) error {
	state := NewGameState(config)
	port := config.Port()
	if port == "" {
		return errors.New("no port set")
	}
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}
	defer ln.Close()
	logger.Info("server started", "port", config.Port())
	for {
		conn, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("accept: %w", err)
		}
		logger.Debug("connection accepted: " + conn.RemoteAddr().String())
		go handle(state, conn, logger)
	}
}

func handle(state *GameState, conn net.Conn, logger *slog.Logger) {
	defer conn.Close()
	player, err := state.AddPlayer(conn)
	if err != nil {
		logger.Error("failed to add player", "error", err)
		protocol.Write(conn, protocol.MessageTypeError, protocol.ErrorPayload{
			Code:    protocol.ErrorCodeServerFull,
			Message: "server is full",
		})
		return
	}
	logPlayerGroup := slog.Group("player", "id", player.ID, "ip_address", player.IPAddress)
	logger.Info("client connected", logPlayerGroup)
	defer state.RemovePlayer(player)
	protocol.Write(conn, protocol.MessageTypeWelcome, protocol.MessagePayloadWelcome{
		PlayerID: player.ID,
	})
	dec := json.NewDecoder(conn)
	for {
		var msg protocol.Message
		if err := dec.Decode(&msg); err != nil {
			if errors.Is(err, io.EOF) {
				logger.Info("client disconnected", logPlayerGroup)
				return
			}
			logger.Warn("bad message received", "error", err, logPlayerGroup)
			protocol.Write(conn, protocol.MessageTypeError, protocol.ErrorPayload{
				Code:    protocol.ErrorCodeBadMessage,
				Message: "bad message received",
			})
			return
		}
		handleClientMessage(msg, logger)
	}
}

func handleClientMessage(msg protocol.Message, logger *slog.Logger) {
	switch msg.Type {
	case protocol.MessageTypeMove:
		var mp protocol.MessagePayloadMove
		err := json.Unmarshal(msg.Payload, &mp)
		if err != nil {
			logger.Error("failed to unmarshal json", "error", err)
		}
	}
}
