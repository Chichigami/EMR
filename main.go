package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/a-h/templ"
	"github.com/chichigami/EMR/internal/database"
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

	cfg := Config{
		db: dbQueries,
	}

	router := gin.Default() //default page should be login page
	router.Static("/static", "../frontend")
	router.LoadHTMLFiles("../frontend/login.html", "../frontend/dashboard.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	router.GET("/patient/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "patient_form.html", gin.H{})
	})
	router.Delims("{[{", "}]}")

	router.GET("/ping", handlerPlaceholder)

	login := router.Group("/login")
	{
		login.POST("", cfg.handlerUsersRead)     //wait for login info, verifies
		login.GET("", cfg.handlerUsersCreate)    //
		login.PUT("", cfg.handlerUsersUpdate)    //update login
		login.DELETE("", cfg.handlerUsersDelete) //delete login, need admin priv
	}
	dashboard := router.Group("/dashboard") //might just make it one handler. cache date in server. if date is today then grab cache?
	{
		dashboard.GET("", cfg.handlerDashboardToday)          //show today's dashboard
		dashboard.GET("/:date", cfg.handlerDashboardNotToday) //show some date's dashboard (maybe yesterday)
	}

	patient := router.Group("/patients")
	{
		patient.POST("/new", cfg.handlerPatientsCreate)   //add new patient
		patient.GET("/:ID", cfg.handlerPatients)          //show patient based on ID
		patient.PUT("/:ID", cfg.handlerPatientsUpdate)    //update patient info
		patient.DELETE("/:ID", cfg.handlerPatientsDelete) //delete existing patient, need admin perm
		//patient.GET("/:ID", handlerPatientQuery)     //query for patient based on patient id, name, dob
	}

	charts := router.Group("/patients/charts")
	{
		charts.POST("/:ID", cfg.handlerChartsCreate)   //make new chart for patient
		charts.GET("/:ID", cfg.handlerChartsRead)      //show chart info
		charts.PUT("/:ID", cfg.handlerChartsUpdate)    //update chart
		charts.DELETE("/:ID", cfg.handlerChartsDelete) //delete chart
	}

	schedule := router.Group("/schedule")
	{
		schedule.POST("/:ID", cfg.handlerAppointmentsCreate)   //schedule a patient
		schedule.DELETE("/:ID", cfg.handlerAppointmentsDelete) //delete a patient's appointment
	}
	router.Run(":8000")
}

type Config struct {
	db *database.Queries
}
