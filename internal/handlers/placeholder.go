package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerPlaceholder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
