package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/a-h/templ"
	"github.com/chichigami/EMR/internal/database"
	"github.com/chichigami/EMR/internal/handlers"
	"github.com/chichigami/EMR/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	dbConnection, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	if err := dbConnection.Ping(); err != nil {
		log.Fatalf("Error pinging database: %s", err)
	}
	defer dbConnection.Close()

	dbQueries := database.New(dbConnection)

	cfg := models.Config{
		Datebase: dbQueries,
	}

	h := handlers.NewHandlerConfig(&cfg)

	router := gin.Default() //default page should be login page

	router.Static("/static", "./internal/static")
	router.LoadHTMLGlob("internal/templates/*")
	//router.Delims("{[{", "}]}")
	//router.GET("/ping", handlerPlaceholder)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	login := router.Group("/login")
	{
		login.GET("", h.HandlerUsersCreate)    //create new login if admin auth
		login.POST("", h.HandlerUsersRead)     //wait for login info, verifies
		login.PUT("", h.HandlerUsersUpdate)    //update login
		login.DELETE("", h.HandlerUsersDelete) //delete login, need admin priv
	}
	dashboard := router.Group("/dashboard") //might just make it one handler. cache date in server. if date is today then grab cache?
	{
		dashboard.GET("", h.HandlerDashboardToday)          //show today's dashboard
		dashboard.GET("/:date", h.HandlerDashboardNotToday) //show some date's dashboard (maybe yesterday)
	}

	patient := router.Group("/patients")
	{
		patient.POST("/new", h.HandlerPatientsCreate)   //add new patient
		patient.GET("/:ID", h.HandlerPatients)          //show patient based on ID
		patient.PUT("/:ID", h.HandlerPatientsUpdate)    //update patient info
		patient.DELETE("/:ID", h.HandlerPatientsDelete) //delete existing patient, need admin perm
		//patient.GET("/:ID", handlerPatientQuery)     //query for patient based on patient id, name, dob
	}

	charts := router.Group("/patients/charts")
	{
		charts.POST("/:ID", h.HandlerChartsCreate)   //make new chart for patient
		charts.GET("/:ID", h.HandlerChartsRead)      //show chart info
		charts.PUT("/:ID", h.HandlerChartsUpdate)    //update chart
		charts.DELETE("/:ID", h.HandlerChartsDelete) //delete chart
	}

	schedule := router.Group("/schedule")
	{
		schedule.POST("/:ID", h.HandlerAppointmentsCreate)   //schedule a patient
		schedule.DELETE("/:ID", h.HandlerAppointmentsDelete) //delete a patient's appointment
	}
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	router.Run(":8000")
}
