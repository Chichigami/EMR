package handlers

import "github.com/chichigami/EMR/internal/models"

type HandlerConfig struct {
	Config *models.Config
}

// makes a config for handlers to connect to database
func NewHandlerConfig(cfg *models.Config) *HandlerConfig {
	return &HandlerConfig{
		Config: cfg,
	}
}
