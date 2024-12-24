package main

import "github.com/gin-gonic/gin"

// func handlerDashboard(c *gin.Context) {
// 	//if date = today, grab cached dashboard
// 	handlerPlaceholder(c)
// }

func (cfg *Config) handlerDashboardToday(c *gin.Context) {
	//run this before work day starts
	handlerPlaceholder(c)
}

func (cfg *Config) handlerDashboardNotToday(c *gin.Context) {
	//run this when date != today
	handlerPlaceholder(c)
}

// type dashboard struct {
// }
