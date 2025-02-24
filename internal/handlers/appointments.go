package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chichigami/EMR/internal/database"
	"github.com/chichigami/EMR/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *HandlerConfig) HandlerAppointmentsCreate(c *gin.Context) {
	var param models.Appointment
	if err := c.ShouldBindJSON(&param); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	dateTime, err := time.Parse("2006-01-02T15:04", param.DateAndTime)
	if err != nil {
		log.Printf("conversion failed")
	}

	patientID, err := strconv.Atoi(param.PatientID)
	if err != nil {
		log.Printf("conversion failed")
	}

	_, err = h.Config.Database.CreateAppointmentForPatient(c, database.CreateAppointmentForPatientParams{
		PatientID: int32(patientID),
		DateOf:    dateTime,
		Reasoning: NullString(param.Reason),
	})
	if err != nil {
		log.Println(err)
	}

	//make this a partial div refresh later on
	c.Header("HX-Refresh", "true")
}

func (h *HandlerConfig) HandlerAppointmentsUpdate(c *gin.Context) {
	param := models.Appointment{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// err := h.Config.Database.UpdateAppointment(c, database.UpdateAppointmentParams{
	// 	ID: ,
	// 	DateOf: ,
	// })
	c.String(200, "updated")
}

func (h *HandlerConfig) HandlerAppointmentsDelete(c *gin.Context) {
	//delete button that sends a delete request with appt id
	id := c.Param("id")
	uuid := uuid.MustParse(id)
	err := h.Config.Database.DeleteAppointment(c, uuid)
	if err != nil {
		log.Println(err.Error())
	}
	c.Header("HX-Location", "")
}
