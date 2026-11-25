package server

import (
	"bufio"
	"errors"
	"fmt"
	"log/slog"
	"net"
)

func StartServer(config ServerConfig, logger *slog.Logger) error {
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
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return
		}
		conn.Write([]byte("din mor: " + line))
	}
}
