package handlers

import "github.com/gin-gonic/gin"

// func handlerDashboard(c *gin.Context) {
// 	//if date = today, grab cached dashboard
// 	handlerPlaceholder(c)
// }

func (cfg *HandlerConfig) HandlerDashboardToday(c *gin.Context) {
	//run this before work day starts
	HandlerPlaceholder(c)
}

func (cfg *HandlerConfig) HandlerDashboardNotToday(c *gin.Context) {
	//run this when date != today
	HandlerPlaceholder(c)
}

// type dashboard struct {
// }
