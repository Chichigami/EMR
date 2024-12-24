package main

import (
	"time"

	"github.com/chichigami/EMR/internal/database"
	"github.com/gin-gonic/gin"
)

func (cfg *Config) handlerAppointmentsNew(c *gin.Context) {
	//
	handlerPlaceholder(c)
}

func (cfg *Config) handlerAppointmentsDelete(c *gin.Context) {

	param := database.DeleteAppointmentParams{
		ChartID:         42069,
		AppointmentDate: time.Date(time.Now().Year(), 12, int(1), 12, 30, 00, 0, time.Now().Location()),
	}
	cfg.db.DeleteAppointment(c.Request.Context(), param)
	handlerPlaceholder(c)
}
