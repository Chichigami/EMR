package handlers

import (
	"fmt"
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
	//think htmx allows for custom json object. can try and format it as a time.Time{}?
	//will look into it afterwards
	param := models.Appointment{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Printf("param: %v, date time: %v\n", param, param.DateAndTime)
	dateTime, err := time.Parse("2006-01-02T15:04", param.DateAndTime)
	fmt.Println(dateTime)
	if err != nil {
		log.Printf("conversion failed")
	}
	patientID, err := strconv.Atoi(param.PatientID)
	if err != nil {
		log.Printf("conversion failed")
	}
	appt, err := h.Config.Database.CreateAppointmentForPatient(c, database.CreateAppointmentForPatientParams{
		PatientID:   int32(patientID),
		Appointment: dateTime,
		Reasoning:   NullStringCheck(param.Reason),
	})
	fmt.Println(appt)
	if err != nil {

	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "appointment created",
		"info":    appt,
	})
}

// func (h *HandlerConfig) HandlerAppointmentsUpdate(c *gin.Context) {
// 	//edit form
// 	param := models.Appointment{}
// 	if err := c.ShouldBindJSON(&param); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	HandlerPlaceholder(c)
// }

func (h *HandlerConfig) HandlerAppointmentsDelete(c *gin.Context) {
	//delete button that sends a delete request with appt id
	id := c.Param("id")
	log.Printf("reached%s\n", id)
	uuid := uuid.MustParse(id)
	err := h.Config.Database.DeleteAppointment(c, uuid)
	if err != nil {
		log.Println(err.Error())
	}
	c.Header("HX-Location", "")
}
