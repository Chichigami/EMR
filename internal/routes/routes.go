package routes

import (
	"net/http"

	"github.com/chichigami/EMR/internal/components"
	"github.com/chichigami/EMR/internal/handlers"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine, h *handlers.HandlerConfig) {
	r.Static("/assets", "./internal/assets")

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/", func(c *gin.Context) {
		page := components.Base("EMR Login", nil, components.Login(), components.DefaultFooter())
		c.Header("Content-Type", "text/html; charset=utf-8")
		if err := page.Render(c, c.Writer); err != nil {
			c.String(http.StatusInternalServerError, "Failed to render page: %v", err)
			return
		}
	})
	login := r.Group("/login")
	{
		login.GET("", h.HandlerUsersCreate)    //create new login if admin auth
		login.POST("", h.HandlerUsersRead)     //wait for login info, verifies
		login.PUT("", h.HandlerUsersUpdate)    //update login
		login.DELETE("", h.HandlerUsersDelete) //delete login, need admin priv
	}
	// dashboard := router.Group("/dashboard") //might just make it one handler. cache date in server. if date is today then grab cache?
	// {
	// 	dashboard.GET("*date", h.HandlerDashboard) //defaults to current day dashboard
	// }

	r.GET("/patients/new", func(c *gin.Context) {
		page := components.Base("New Patient", components.DefaultNavbar(), components.PatientForm(), components.DefaultFooter())
		c.Header("Content-Type", "text/html; charset=utf-8")
		if err := page.Render(c, c.Writer); err != nil {
			c.String(http.StatusInternalServerError, "Failed to render page: %v", err)
			return
		}
	})

	patient := r.Group("/patients")
	{
		patient.POST("/create", h.HandlerPatientsCreate) //add new patient
		patient.GET("/:ID", h.HandlerPatientsRead)       //show patient based on ID
		//patient.PUT("/:ID", h.HandlerPatientsUpdate)     //update patient info
		patient.DELETE("/:ID", h.HandlerPatientsDelete) //delete existing patient, need admin perm
		//patient.GET("/:ID", handlerPatientQuery)     //query for patient based on patient id, name, dob
	}

	// charts := r.Group("/patients/charts")
	// {
	// 	charts.POST("/:ID", h.HandlerChartsCreate)   //make new chart for patient
	// 	charts.GET("/:ID", h.HandlerChartsRead)      //show chart info
	// 	charts.PUT("/:ID", h.HandlerChartsUpdate)    //update chart
	// 	charts.DELETE("/:ID", h.HandlerChartsDelete) //delete chart
	// }

	r.GET("/schedule", func(c *gin.Context) {
		page := components.Base("New Appointment", components.DefaultNavbar(), components.Appointment(), components.DefaultFooter())
		c.Header("Content-Type", "text/html; charset=utf-8")
		if err := page.Render(c, c.Writer); err != nil {
			c.String(http.StatusInternalServerError, "Failed to render page: %v", err)
			return
		}
	})
	schedule := r.Group("/schedule")
	{
		schedule.POST("/create", h.HandlerAppointmentsCreate) //schedule a patient
		//schedule.DELETE("/:ID", h.HandlerAppointmentsDelete) //delete a patient's appointment
	}

	r.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
