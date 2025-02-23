package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chichigami/EMR/internal/components"
	"github.com/chichigami/EMR/internal/database"
	"github.com/chichigami/EMR/internal/models"
	"github.com/gin-gonic/gin"
)

// get patient data and show dashboard of charts, appointments, etc
//
// GET
func (h *HandlerConfig) HandlerPatientsRead(c *gin.Context) {
	id := c.Param("id")
	patientID, err := ConvertStringToInt32(id)
	dbPatient, err := h.Config.Database.GetPatient(c, patientID)
	if errors.Is(err, sql.ErrNoRows) {
		c.String(http.StatusOK, "Patient does not exist")
		return
	} else if err != nil {
		log.Printf("get patient error: %s", err.Error())
		return
	}

	//go routine to fetch patient appointments and charts
	//if err is no row then ignore
	dbPatientAppt, err := h.Config.Database.GetAppointmentBasedOnPatient(c, int32(patientID))
	dbPatientChart, err := h.Config.Database.GetPatientCharts(c, int32(patientID))

	patientData := models.PatientProfile{
		Demographic:  dbPatient,
		ID:           id,
		Appointments: dbPatientAppt,
		Charts:       dbPatientChart,
	}
	h.renderPatientDashbord(c, patientData)
}

// makes a new patient
//
// POST
func (h *HandlerConfig) HandlerPatientsCreate(c *gin.Context) {
	param := models.Patient{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	date, err := time.Parse("2006-01-02", param.DateOfBirth)
	if err != nil {
		log.Printf("Failed to convert date: input=%s, error=%v", param.DateOfBirth, err)
		c.String(http.StatusInternalServerError, "converting date/time failed")
		return
	}
	patient, err := h.Config.Database.CreatePatient(c, database.CreatePatientParams{
		FirstName:            param.FirstName,
		MiddleName:           NullString(param.MiddleName),
		LastName:             param.LastName,
		DateOfBirth:          date,
		Sex:                  param.Sex,
		Gender:               param.Gender,
		SocialSecurityNumber: NullString(param.SocialSecurityNumber),
		Pharmacy:             param.Pharmacy,
		Email:                NullString(param.Email),
		LocationAddress:      param.LocationAddress,
		ZipCode:              param.ZipCode,
		CellPhoneNumber:      NullString(param.CellPhoneNumber),
		HomePhoneNumber:      NullString(param.HomePhoneNumber),
		MaritalStatus:        NullString(param.MaritalStatus),
		Insurance:            NullString(param.Insurance),
		PrimaryCareDoctor:    NullString(param.PrimaryCareDoctor),
		ExtraNote:            NullString(param.ExtraNotes),
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "making new patient failed")
		return
	}

	//redirect to patient page? or something else
	c.JSON(http.StatusCreated, gin.H{
		"message":    "success",
		"chart info": patient,
	})
}

func (h *HandlerConfig) HandlerPatientsDelete(c *gin.Context) {
	id := c.Param("id")
	patientID, err := ConvertStringToInt32(id)
	if err != nil {
		log.Println(err.Error())
	}
	err = h.Config.Database.DeletePatient(c, patientID)
	if err != nil {
		log.Printf("could not delete patient: %v\n", patientID)
		c.String(http.StatusInternalServerError, "failed to delete patient")
		return
	}

	c.Header("HX-Redirect", "/patients/deleted")
}

func (h *HandlerConfig) HandlerPatientDeleteAll(c *gin.Context) {
	err := h.Config.Database.DeleteAllPatients(c)
	if err != nil {
		c.String(http.StatusOK, "DELETED ALL PATIENTS")
	}
}

func HandlerPatientDNE(c *gin.Context) {
	c.String(http.StatusOK, "Patient does not exist")
}

func (h *HandlerConfig) renderPatientDashbord(c *gin.Context, patient models.PatientProfile) {
	page := components.Base(fmt.Sprintf("%s, %s (%s)'s Profile", patient.Demographic.LastName, patient.Demographic.FirstName, patient.ID),
		components.PatientNavbar(patient),
		components.PatientDashboard(patient),
		components.DefaultFooter())
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := page.Render(c, c.Writer); err != nil {
		c.String(http.StatusInternalServerError, "Failed to render page: %v", err)
		return
	}
}
