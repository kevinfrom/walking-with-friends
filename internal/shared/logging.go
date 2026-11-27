package shared

import (
	"log/slog"
	"os"
	"strconv"
)

func NewLogger(config BaseConfig) *slog.Logger {
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	logLevel := slog.LevelInfo
	if debug {
		logLevel = slog.LevelDebug
	}
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel,
	})
	return slog.New(handler)
}
