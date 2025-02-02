package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/chichigami/EMR/internal/database"
	"github.com/chichigami/EMR/internal/models"
	"github.com/gin-gonic/gin"
)

func ConvertStringToDate(s string) (time.Time, error) {
	//string to time.Time YYYYMMDD
	//if s contains delimiters then split otherwise don't
	splitted := strings.Split(s, "-")
	formatted := fmt.Sprintf("%s%s%s", splitted[0], splitted[1], splitted[2])
	fmt.Println(formatted)
	if len(formatted) != 8 {
		return time.Time{}, fmt.Errorf("need 8 numbers in YYYYMMDD format")
	}
	const shortForm = "20060102"
	date, err := time.Parse(shortForm, formatted)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

// get patient data
//
// GET
func (h *HandlerConfig) HandlerPatientsRead(c *gin.Context) {
	patientStr := c.Param(":id")
	patientID32, err := strconv.ParseInt(patientStr, 10, 32)
	if err != nil {
		log.Fatalf("converting patient id string to int failed")
	}
	dbPatient, err := h.Config.Datebase.GetPatient(c, database.NullInt32(int32(patientID32)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(400, gin.H{
		"patient": dbPatient,
	})
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

	date, err := ConvertStringToDate(param.DateOfBirth)
	if err != nil {
		log.Printf("Failed to convert date: input=%s, error=%v", param.DateOfBirth, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "converting date error",
		})
		return
	}
	patient, err := h.Config.Datebase.CreatePatient(c, database.CreatePatientParams{
		FirstName:            param.FirstName,
		MiddleName:           database.NullStringCheck(param.MiddleName),
		LastName:             param.LastName,
		DateOfBirth:          date,
		Sex:                  param.Sex,
		Gender:               param.Gender,
		SocialSecurityNumber: database.NullStringCheck(param.SocialSecurityNumber),
		Pharmacy:             param.Pharmacy,
		Email:                database.NullStringCheck(param.Email),
		LocationAddress:      param.LocationAddress,
		ZipCode:              param.ZipCode,
		CellPhoneNumber:      database.NullStringCheck(param.CellPhoneNumber),
		HomePhoneNumber:      database.NullStringCheck(param.HomePhoneNumber),
		MaritalStatus:        database.NullStringCheck(param.MaritalStatus),
		Insurance:            database.NullStringCheck(param.Insurance),
		PrimaryCareDoctor:    database.NullStringCheck(param.PrimaryCareDoctor),
		ExtraNote:            database.NullStringCheck(param.ExtraNotes),
	})
	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":    "success",
		"chart info": patient,
	})
}

func (h *HandlerConfig) HandlerPatientsDelete(c *gin.Context) {
	err := h.Config.Datebase.DeleteAllPatients(c)
	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "patient reset",
	})
}

func (h *HandlerConfig) HandlerPatientsUpdate(c *gin.Context) {
	//get patient data
	//check which data collides
	//update data
	HandlerPlaceholder(c)
}
