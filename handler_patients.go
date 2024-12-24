package main

import (
	"github.com/gin-gonic/gin"
)

func (cfg *Config) handlerPatients(c *gin.Context) {
	//get patient info
	handlerPlaceholder(c)
}

func (cfg *Config) handlerPatientsNew(c *gin.Context) {
	//make new patient into db
	handlerPlaceholder(c)
}

func (cfg *Config) handlerPatientsDelete(c *gin.Context) {
	//delete patient from db
	handlerPlaceholder(c)
}

// func handlerPatientQuery(c *gin.Context) {
// 	//client is unsure who patient is. query via id, name, dob
// 	//call handlerPatient after
// 	handlerPlaceholder(c)
// }

func (cfg *Config) handlerPatientsUpdate(c *gin.Context) {
	handlerPlaceholder(c)
}
