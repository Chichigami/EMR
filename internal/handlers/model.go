package handlers

import "github.com/chichigami/EMR/internal/models"

type HandlerConfig struct {
	Config *models.Config
}

func NewHandlerConfig(cfg *models.Config) *HandlerConfig {
	return &HandlerConfig{
		Config: cfg,
	}
}
