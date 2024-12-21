package main

import "github.com/gin-gonic/gin"

func handlerChartNew(c *gin.Context) {
	//wait for request from front end. make new chart.
	handlerPlaceholder(c)
}

func handlerChart(c *gin.Context) {
	//get chart based on ID
	handlerPlaceholder(c)
}

func handlerChartUpdate(c *gin.Context) {
	//should receive a request every chart completion? or every minute.
	handlerPlaceholder(c)
}

func handlerChartDelete(c *gin.Context) {
	//delete chart based on id
	handlerPlaceholder(c)
}
