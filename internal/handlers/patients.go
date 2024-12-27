package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *HandlerConfig) HandlerPatients(c *gin.Context) {
	HandlerPlaceholder(c)
}

func (h *HandlerConfig) HandlerPatientsCreate(c *gin.Context) {
	//make new patient into db
	HandlerPlaceholder(c)
}

func (h *HandlerConfig) HandlerPatientsDelete(c *gin.Context) {
	//delete patient from db
	HandlerPlaceholder(c)
}

// func handlerPatientQuery(c *gin.Context) {
// 	//client is unsure who patient is. query via id, name, dob
// 	//call handlerPatient after
// 	handlerPlaceholder(c)
// }

func (h *HandlerConfig) HandlerPatientsUpdate(c *gin.Context) {
	HandlerPlaceholder(c)
}
