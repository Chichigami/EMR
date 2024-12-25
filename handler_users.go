package main

import (
	"log"
	"net/http"

	"github.com/chichigami/EMR/internal/database"
	"github.com/gin-gonic/gin"
)

func (cfg *Config) handlerUsersCreate(c *gin.Context) {
	//check auth, if not admin, then return error
	//from a form
	//hash password
	account := database.CreateUserParams{
		Username:       "",
		HashedPassword: "",
		LastName:       "",
		FirstName:      "",
		Permissions:    "",
	}
	_, err := cfg.db.CreateUser(c.Request.Context(), account)
	if err != nil {
		log.Fatalf("User creation failed: %s", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
	})
}

func (cfg *Config) handlerUsersRead(c *gin.Context) {
	//get info from form
	//hash the password
	//get password from db using username
	//compare
	handlerPlaceholder(c)
}

func (cfg *Config) handlerUsersUpdate(c *gin.Context) {
	//check if auth matches auth in db or is admin
	//update password if either
	handlerPlaceholder(c)
}

func (cfg *Config) handlerUsersDelete(c *gin.Context) {
	//check if auth is admin
	//if not then return error
	//delete user
	handlerPlaceholder(c)
}
