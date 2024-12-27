package Handlers

import (
	"time"

	"github.com/chichigami/EMR/internal/database"
	"github.com/gin-gonic/gin"
)

func (h *HandlerConfig) HandlerAppointmentsCreate(c *gin.Context) {
	//
	HandlerPlaceholder(c)
}

func (h *HandlerConfig) HandlerAppointmentsUpdate(c *gin.Context) {
	//do a table update

	HandlerPlaceholder(c)
}

func (h *HandlerConfig) HandlerAppointmentsDelete(c *gin.Context) {

	param := database.DeleteAppointmentParams{
		ChartID:         42069,
		AppointmentDate: time.Date(time.Now().Year(), 12, int(1), 12, 30, 00, 0, time.Now().Location()),
	}
	h.Config.Datebase.DeleteAppointment(c, param)
	HandlerPlaceholder(c)
}
