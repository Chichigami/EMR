package handlers

import "github.com/gin-gonic/gin"

func (h *HandlerConfig) HandlerChartsCreate(c *gin.Context) {
	//wait for request from front end. make new chart.
	HandlerPlaceholder(c)
}

func (h *HandlerConfig) HandlerChartsRead(c *gin.Context) {
	//get chart based on ID
	HandlerPlaceholder(c)
}

func (h *HandlerConfig) HandlerChartsUpdate(c *gin.Context) {
	//should receive a request every chart completion? or every minute.
	HandlerPlaceholder(c)
}

func (h *HandlerConfig) HandlerChartsDelete(c *gin.Context) {
	//delete chart based on id
	HandlerPlaceholder(c)
}
