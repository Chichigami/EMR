package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//connect to DB

	r := gin.Default() //default page should be login page

	r.GET("/ping", handlerPlaceholder)

	login := r.Group("/login")
	{
		login.POST("", handlerLoginNew)      //new login, need admin priv
		login.GET("", handlerLogin)          //wait for login info, verifies
		login.PUT("", handlerLoginUpdate)    //update login
		login.DELETE("", handlerLoginDelete) //delete login, need admin priv
	}
	dashboard := r.Group("dashboard") //might just make it one handler. cache date in server. if date is today then grab cache?
	{
		dashboard.GET("", handlerDashboardToday)          //show today's dashboard
		dashboard.GET("/:date", handlerDashboardNotToday) //show some date's dashboard (maybe yesterday)
	}

	patient := r.Group("patients")
	{
		patient.POST("", handlerPatientNew)          //add new patient
		patient.GET("/:ID", handlerPatient)          //show patient based on ID
		patient.PUT("/:ID", handlerPatientUpdate)    //update patient info
		patient.DELETE("/:ID", handlerPatientDelete) //delete existing patient, need admin perm
		//patient.GET("/:ID", handlerPatientQuery)     //query for patient based on patient id, name, dob
	}

	charts := r.Group("charts")
	{
		charts.POST("/:ID", handlerChartNew)      //make new chart for patient
		charts.GET("/:ID", handlerChart)          //show chart info
		charts.PUT("/:ID", handlerChartUpdate)    //update chart
		charts.DELETE("/:ID", handlerChartDelete) //delete chart
	}

	schedule := r.Group("/schedule")
	{
		schedule.POST("/:ID", handlerScheduleNew)      //schedule a patient
		schedule.DELETE("/:ID", handlerScheduleDelete) //delete a patient's appointment
	}
	r.Run(":8000")
}
