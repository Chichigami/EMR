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
	r.GET("/", renderLoginView)
	r.GET("/patients/new", renderCreatePatientView)
	r.GET("/schedule", renderScheduleView)
	r.GET("/frontdesk", renderFrontDeskView)
	r.GET("/test", renderTestView)

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
		dashboard.GET("", h.HandlerDashboard)
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
		patients.GET("/deleted", renderDeleted)
	}

	charts := patients.Group("/:id/charts")
	{
		charts.POST("/new", h.HandlerChartsCreate) //make new chart for patient
		charts.GET("/:uuid", h.HandlerChartsGet)   //show chart info
		// charts.PUT("/:id", h.HandlerChartsUpdate)    //update chart
		// charts.DELETE("/:id", h.HandlerChartsDelete) //delete chart
	}

	schedule := r.Group("/schedule")
	{
		schedule.GET("/modal/:id", renderAppointmentModal)
		schedule.PUT("/:id", h.HandlerAppointmentsUpdate)
		schedule.POST("/create", h.HandlerAppointmentsCreate) //schedule a patient
		schedule.DELETE("/:id", h.HandlerAppointmentsDelete)  //delete a patient's appointment
		schedule.GET("/null", nullModal)
	}
}

func renderTestView(c *gin.Context) {
	page := components.Base("New Patient", nil, nil, components.ChartFooter("2138192012314-412412481243-123123", "dr. feng"))
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := page.Render(c, c.Writer); err != nil {
		c.String(http.StatusInternalServerError, "Failed to render page: %v", err)
		return
	}
}

func renderCreatePatientView(c *gin.Context) {
	page := components.Base("New Patient", components.DefaultNavbar(), components.PatientForm(), components.DefaultFooter())
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := page.Render(c, c.Writer); err != nil {
		c.String(http.StatusInternalServerError, "Failed to render page: %v", err)
		return
	}
}

func renderFrontDeskView(c *gin.Context) {
	page := components.Base("EMR Frontdesk", components.DefaultNavbar(), nil, components.DefaultFooter())
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := page.Render(c, c.Writer); err != nil {
		c.String(http.StatusInternalServerError, "Failed to render page: %v", err)
		return
	}
}

func renderLoginView(c *gin.Context) {
	page := components.Base("EMR Login", nil, components.Login(), components.DefaultFooter())
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := page.Render(c, c.Writer); err != nil {
		c.String(http.StatusInternalServerError, "Failed to render page: %v", err)
		return
	}
}

func renderScheduleView(c *gin.Context) {
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

func renderDeleted(c *gin.Context) {
	c.String(200, "Patient deleted")
}

func nullModal(c *gin.Context) {
	c.String(http.StatusOK, "")
}
