package main

import "github.com/gin-gonic/gin"

func handlerPatient(c *gin.Context) {
	//get patient info
	handlerPlaceholder(c)
}

func handlerPatientNew(c *gin.Context) {
	//make new patient into db
	handlerPlaceholder(c)
}

func handlerPatientDelete(c *gin.Context) {
	//delete patient from db
	handlerPlaceholder(c)
}

// func handlerPatientQuery(c *gin.Context) {
// 	//client is unsure who patient is. query via id, name, dob
// 	//call handlerPatient after
// 	handlerPlaceholder(c)
// }

func handlerPatientUpdate(c *gin.Context) {
	handlerPlaceholder(c)
}
