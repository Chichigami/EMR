package handlers

import (
	"log"
	"net/http"

	"github.com/chichigami/EMR/internal/auth"
	"github.com/chichigami/EMR/internal/database"
	"github.com/chichigami/EMR/internal/models"
	"github.com/gin-gonic/gin"
)

// check auth, if not admin, then return error
// from a form
// hash password
func (h *HandlerConfig) HandlerUsersCreate(c *gin.Context) {
	param := models.UserLogin{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//insert auth check
	hashed, err := auth.HashPassword(param.Password)
	account := database.CreateUserParams{
		Username:       param.Username,
		HashedPassword: hashed,
		LastName:       param.Lastname,
		FirstName:      param.Firstname,
		Permissions:    param.Permission,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "trouble making password",
		})
		return
	}
	if err := h.Config.Datebase.CreateUser(c.Request.Context(), account); err != nil {
		log.Fatalf("User creation failed: %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
		"account": param,
	})
	return
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

	account, err := h.Config.Datebase.GetHashedPassword(c, param.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to verify username",
		})
		return
	}

	err = auth.CheckPasswordHash(account.HashedPassword, param.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong password",
		})
	}
	c.Header("HX-Redirect", "/dashboard")
}

// edit button request
//
// updates a user's info
func (h *HandlerConfig) HandlerUsersUpdate(c *gin.Context) {
	param := models.UpdateUserLogin{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	hashed, err := auth.HashPassword(param.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.Config.Datebase.UpdateUserInfo(c, database.UpdateUserInfoParams{
		Username:       "get username from auth from cookie",
		HashedPassword: hashed,
		LastName:       param.Lastname,
		FirstName:      param.Firstname,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "account updated",
	})
	return
}

// button request in account page
//
// deletes account
func (h *HandlerConfig) HandlerUsersDelete(c *gin.Context) {
	//check account owner?
	//hx-confirm
	//hx-delete
	err := h.Config.Datebase.DeleteUser(c, "auth_connection_to_username")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "account deleted",
	})
	return
}
