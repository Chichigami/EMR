package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chichigami/EMR/internal/auth"
	"github.com/chichigami/EMR/internal/database"
	"github.com/chichigami/EMR/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *HandlerConfig) HandlerUsersCreate(c *gin.Context) {
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
	_, err := h.Config.Datebase.CreateUser(c.Request.Context(), account)
	if err != nil {
		log.Fatalf("User creation failed: %s", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
	})
}

// receives a http POST request and verifies login
//
// gives auth token
func (h *HandlerConfig) HandlerUsersRead(c *gin.Context) {
	param := models.UserLogin{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Printf("user: %s \npass: %s\n", param.Username, param.Password)
	hashed, err := auth.HashPassword(param.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password hashing failed",
		})
		return
	}
	err = auth.CheckPasswordHash(hashed, param.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong password",
		})
	}
	c.Header("HX-Redirect", "/dashboard")
}

func (h *HandlerConfig) HandlerUsersUpdate(c *gin.Context) {
	//check if auth matches auth in db or is admin
	//update password if either
	HandlerPlaceholder(c)
}

func (h *HandlerConfig) HandlerUsersDelete(c *gin.Context) {
	//check if auth is admin
	//if not then return error
	//delete user
	HandlerPlaceholder(c)
}

//c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
