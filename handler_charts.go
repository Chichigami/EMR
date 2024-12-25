package main

import "github.com/gin-gonic/gin"

func (cfg *Config) handlerChartsCreate(c *gin.Context) {
	//wait for request from front end. make new chart.
	handlerPlaceholder(c)
}

func (cfg *Config) handlerChartsRead(c *gin.Context) {
	//get chart based on ID
	handlerPlaceholder(c)
}

func (cfg *Config) handlerChartsUpdate(c *gin.Context) {
	//should receive a request every chart completion? or every minute.
	handlerPlaceholder(c)
}

func (cfg *Config) handlerChartsDelete(c *gin.Context) {
	//delete chart based on id
	handlerPlaceholder(c)
}
