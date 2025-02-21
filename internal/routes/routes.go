package routes

import (
	"net/http"
	"time"

	"github.com/chichigami/EMR/internal/components"
	"github.com/chichigami/EMR/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine, h *handlers.HandlerConfig) {
	r.Static("/assets", "./internal/assets")
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/graceful", func(c *gin.Context) {
		c.String(http.StatusOK, "Ping received! Processing...")
		time.Sleep(2 * time.Second)
		c.String(http.StatusOK, "Ping completed!")
	})

	//start of views
	r.GET("/", renderLoginPage)
	r.GET("/patients/new", renderCreatePatientPage)
	r.GET("/schedule", renderSchedulePage)

	//start of api endpoints
	login := r.Group("/login")
	{
		login.GET("", h.HandlerUsersCreate) //create new login if admin auth
		login.POST("", h.HandlerUsersRead)  //wait for login info, verifies
		login.PUT("", h.HandlerUsersUpdate) //update login
		//login.DELETE("", h.HandlerUsersDelete) //delete login, need admin priv
		login.DELETE("", h.HandlerDeleteAllUser)
	}

	dashboard := r.Group("/dashboard")
	{
		dashboard.GET("", h.HandlerDashboard) //defaults to current day dashboard
		// 	dashboard.POST("*date/items", h.HandlerDashboardUpdate)
	}

	patients := r.Group("/patients")
	{
		patients.POST("/create", h.HandlerPatientsCreate) //add new patient
		patients.GET("/:id", h.HandlerPatientsRead)       //show patient based on id
		//patients.PUT("/:id", h.HandlerPatientsUpdate)     //update patient info
		patients.DELETE("/:id", h.HandlerPatientsDelete) //delete existing patient, need admin perm
		//patients.GET("/:id", handlerPatientQuery)     //query for patient based on patient id, name, dob
		patients.GET("/dne", handlers.HandlerPatientDNE)
		patients.POST("/danger/delete", h.HandlerPatientDeleteAll)
	}

	// charts := r.Group("/patients/charts")
	// {
	// 	charts.POST("/:id", h.HandlerChartsCreate)   //make new chart for patient
	// 	charts.GET("/:id", h.HandlerChartsRead)      //show chart info
	// 	charts.PUT("/:id", h.HandlerChartsUpdate)    //update chart
	// 	charts.DELETE("/:id", h.HandlerChartsDelete) //delete chart
	// }

	schedule := r.Group("/schedule")
	{
		schedule.GET("/modal/:id", renderAppointmentModal)
		schedule.POST("/create", h.HandlerAppointmentsCreate) //schedule a patient
		schedule.DELETE("/:id", h.HandlerAppointmentsDelete)  //delete a patient's appointment
		schedule.GET("/null", nullModal)
	}
}

func renderCreatePatientPage(c *gin.Context) {
	page := components.Base("New Patient", components.DefaultNavbar(), components.PatientForm(), components.DefaultFooter())
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := page.Render(c, c.Writer); err != nil {
		c.String(http.StatusInternalServerError, "Failed to render page: %v", err)
		return
	}
}

func renderLoginPage(c *gin.Context) {
	page := components.Base("EMR Login", nil, components.Login(), components.DefaultFooter())
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := page.Render(c, c.Writer); err != nil {
		c.String(http.StatusInternalServerError, "Failed to render page: %v", err)
		return
	}
}

func renderSchedulePage(c *gin.Context) {
	page := components.Base("New Appointment", components.DefaultNavbar(), components.Appointment(), components.DefaultFooter())
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := page.Render(c, c.Writer); err != nil {
		c.String(http.StatusInternalServerError, "Failed to render page: %v", err)
		return
	}
}

func renderAppointmentModal(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := components.ModalAppointment(c.Param("id")).Render(c, c.Writer); err != nil {
		c.String(http.StatusInternalServerError, "Failed to render page: %v", err)
		return
	}
}

func nullModal(c *gin.Context) {
	c.String(http.StatusOK, "")
}
