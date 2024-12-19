package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//connect to DB

	router := gin.Default() //default page should be login page
	router.Delims("{[{", "}]}")

	router.GET("/ping", handlerPlaceholder)

	login := router.Group("/login")
	{
		login.POST("", handlerLoginNew)      //new login, need admin priv
		login.GET("", handlerLogin)          //wait for login info, verifies
		login.PUT("", handlerLoginUpdate)    //update login
		login.DELETE("", handlerLoginDelete) //delete login, need admin priv
	}
	dashboard := router.Group("/dashboard") //might just make it one handler. cache date in server. if date is today then grab cache?
	{
		dashboard.GET("", handlerDashboardToday)          //show today's dashboard
		dashboard.GET("/:date", handlerDashboardNotToday) //show some date's dashboard (maybe yesterday)
	}

	patient := router.Group("/patients")
	{
		patient.POST("", handlerPatientNew)          //add new patient
		patient.GET("/:ID", handlerPatient)          //show patient based on ID
		patient.PUT("/:ID", handlerPatientUpdate)    //update patient info
		patient.DELETE("/:ID", handlerPatientDelete) //delete existing patient, need admin perm
		//patient.GET("/:ID", handlerPatientQuery)     //query for patient based on patient id, name, dob
	}

	charts := router.Group("/patients/charts")
	{
		charts.POST("/:ID", handlerChartNew)      //make new chart for patient
		charts.GET("/:ID", handlerChart)          //show chart info
		charts.PUT("/:ID", handlerChartUpdate)    //update chart
		charts.DELETE("/:ID", handlerChartDelete) //delete chart
	}

	schedule := router.Group("/schedule")
	{
		schedule.POST("/:ID", handlerScheduleNew)      //schedule a patient
		schedule.DELETE("/:ID", handlerScheduleDelete) //delete a patient's appointment
	}
	router.Run(":8000")
}

type config struct {
	db string
}
