package routes

import (
	"fmt"
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

	r.GET("/", func(c *gin.Context) { //default view => login page
		handlers.RenderView(c, "EMR Login", nil, components.Login(), components.DefaultFooter())
	})

	r.GET("/test", func(c *gin.Context) {
		handlers.RenderView(c,
			"test",
			components.DefaultNavbar(time.Now().Format(time.DateOnly)),
			components.Test_Dashboard(time.Now().Format(time.DateOnly)),
			components.DefaultFooter())
	})

	user := r.Group("/user")
	{
		user.GET("", h.HandlerUsersCreate)
		user.POST("", h.HandlerUsersRead)
		user.PUT("", h.HandlerUsersUpdate)
		//login.DELETE("", h.HandlerUsersDelete) //delete login, need admin priv
		user.DELETE("", h.HandlerDeleteAllUser)
		user.GET("menu", h.RenderUserMenu)
		user.POST("/logout", handlers.HandlerUserLogout)
	}

	dashboard := r.Group("/dashboard")
	{
		dashboard.GET("", h.HandlerDashboard)
		dashboard.PUT("/test", func(ctx *gin.Context) {
			type Moved struct {
				ID   string
				From string
				To   string
			}
			var move Moved
			if err := ctx.ShouldBindJSON(&move); err != nil {
				fmt.Println(move)
			}
		})
		// 	dashboard.POST("*date/items", h.HandlerDashboardUpdate)
	}

	patients := r.Group("/patients")
	{
		patients.GET("/new", func(c *gin.Context) {
			handlers.RenderView(c, "New Patient", components.DefaultNavbar(time.Now().Format(time.DateOnly)), components.PatientForm(), components.DefaultFooter())
		})
		patients.POST("/create", h.HandlerPatientsCreate)
		patients.GET("/:id", h.HandlerPatientsRead)
		//patients.PUT("/:id", h.HandlerPatientsUpdate)
		patients.DELETE("/:id", h.HandlerPatientsDelete)
		//patients.GET("/:id", handlerPatientQuery)
		patients.GET("/dne", handlers.HandlerPatientDNE)
		patients.POST("/danger/delete", h.HandlerPatientDeleteAll)
		patients.GET("/deleted", handlers.RenderDeletedView)
	}

	charts := patients.Group("/:id/charts")
	{
		charts.POST("/new", h.HandlerChartsCreate)
		charts.GET("/:uuid", h.HandlerChartsGet)
		charts.PUT("/:uuid", h.HandlerChartsUpdate)
		charts.DELETE("/:uuid", h.HandlerChartsDelete)
	}

	schedule := r.Group("/schedule")
	{
		schedule.GET("", func(c *gin.Context) {
			handlers.RenderView(c, "New Appointment", components.DefaultNavbar(time.Now().Format("2006-01-02")), components.Appointment(), components.DefaultFooter())
		})
		schedule.GET("/modal/:id", handlers.RenderAppointmentModal)
		schedule.PUT("/:id", h.HandlerAppointmentsUpdate)
		schedule.POST("/create", h.HandlerAppointmentsCreate)
		schedule.DELETE("/:id", h.HandlerAppointmentsDelete)
		schedule.GET("/null", func(c *gin.Context) { //there should be another way to do this.
			c.String(http.StatusOK, "")
		})
	}

	healthz := r.Group("/healthz")
	{
		healthz.GET("", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})
		healthz.GET("/graceful", func(c *gin.Context) {
			c.String(http.StatusOK, "Ping received! Processing...")
			time.Sleep(2 * time.Second)
			c.String(http.StatusOK, "Ping completed!")
		})
	}
}
