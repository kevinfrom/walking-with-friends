package shared

import (
	"os"
	"strconv"
)

type BaseConfig struct {
	debug bool
}

func NewBaseConfig() BaseConfig {
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	return BaseConfig{
		debug: debug,
	}
}

func (c BaseConfig) Debug() bool {
	return c.debug
}
