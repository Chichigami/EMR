package main

import "github.com/gin-gonic/gin"

func handlerLoginNew(c *gin.Context) {
	//make new login, admin perm
	handlerPlaceholder(c)
}

func handlerLogin(c *gin.Context) {
	//GET a req with user, password, server
	//validate user, pass, server
	//response
	//on success redirect
	//on failure show error message
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
