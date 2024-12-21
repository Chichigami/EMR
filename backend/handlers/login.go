package main

import (
	"net/http"

	"github.com/chichigami/EMR/internal/auth"
	"github.com/gin-gonic/gin"
)

func handlerLoginNew(c *gin.Context) {
	//make new login, admin perm
	handlerPlaceholder(c)
}

func handlerLoginVerify(c *gin.Context) {
	//from POST req
	password := "get password from req"
	if hashed, ok := auth.HashPassword(password); ok == nil {
		if success := auth.CheckPasswordHash(hashed, password); success == nil {
			//give auth cookie
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Login successful",
			})
		}
	}
	handlerPlaceholder(c)
}

func handlerLoginUpdate(c *gin.Context) {
	//after verified user, can update user or password. if admin can update other perms.
	handlerPlaceholder(c)
}

func handlerLoginDelete(c *gin.Context) {
	//delete login, admin perm
	handlerPlaceholder(c)
}
