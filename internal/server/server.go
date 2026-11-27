package server

import (
	"bufio"
	"errors"
	"fmt"
	"log/slog"
	"net"
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
		conn.Write([]byte("error: " + err.Error()))
		logger.Error("failed to add player", "error", err)
		return
	}
	logger.Debug("player added", "player id", player.ID)
	defer state.RemovePlayer(player)
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			logger.Error("failed to read line", "error", err)
			return
		}
		logger.Debug("received line", "line", line)
		conn.Write([]byte("echo: " + line))
	}
}
