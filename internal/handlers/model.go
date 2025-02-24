package handlers

import (
	"github.com/chichigami/EMR/internal/config"
)

type HandlerConfig struct {
	Config *config.Config
}

// makes a config for handlers to connect to database
func NewHandlerConfig(cfg *config.Config) *HandlerConfig {
	return &HandlerConfig{
		Config: cfg,
	}
}
